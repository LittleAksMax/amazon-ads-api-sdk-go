package amazon_ads_api_go_sdk

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

const reportingV3ContentType = "application/vnd.createasyncreportrequest.v3+json"

// 148d8c91-a8c6-49c8-bc2f-7d4a0e24e8f4

type ReportsService service

type Report struct {
	service   *ReportsService
	profileID int64
	mu        sync.RWMutex
	details   models.ReportDetails
}

func (s *ReportsService) RequestReport(ctx context.Context, profileID int64, options *models.RequestReportOptions) (*Report, error) {
	if options == nil {
		return nil, errors.New("report options are nil")
	}

	if err := s.client.setToken(); err != nil {
		return nil, err
	}

	req, err := buildJSONRequest(ctx, http.MethodPost, s.client.regionURL(), "reporting/reports", profileID, options, s.client)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", reportingV3ContentType)

	details, err := s.doReportDetailsRequest(req)
	if err != nil {
		return nil, err
	}

	return s.newReport(profileID, details), nil
}

func (s *ReportsService) CancelReport(ctx context.Context, profileID int64, reportID string) error {
	if err := s.client.setToken(); err != nil {
		return err
	}

	req, err := buildJSONRequest(ctx, http.MethodDelete, s.client.regionURL(), fmt.Sprintf("reporting/reports/%s", reportID), profileID, nil, s.client)
	if err != nil {
		return err
	}

	_, err = s.doRequest(req)
	return err
}

func (s *ReportsService) GetReport(ctx context.Context, profileID int64, reportID string) (*Report, error) {
	details, err := s.getReportDetails(ctx, profileID, reportID)
	if err != nil {
		return nil, err
	}

	return s.newReport(profileID, details), nil
}

func (s *ReportsService) newReport(profileID int64, details *models.ReportDetails) *Report {
	return &Report{
		service:   s,
		profileID: profileID,
		details:   *details,
	}
}

func (s *ReportsService) getReportDetails(ctx context.Context, profileID int64, reportID string) (*models.ReportDetails, error) {
	if reportID == "" {
		return nil, errors.New("reportID is empty")
	}

	if err := s.client.setToken(); err != nil {
		return nil, err
	}

	u := url.URL{
		Scheme: "https",
		Host:   s.client.regionURL(),
		Path:   "reporting/reports/" + reportID,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	s.client.setRequestHeaders(req, map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 reportingV3ContentType,
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	})

	return s.doReportDetailsRequest(req)
}

func (s *ReportsService) doReportDetailsRequest(req *http.Request) (*models.ReportDetails, error) {
	bodyBytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	var details models.ReportDetails
	if err := json.Unmarshal(bodyBytes, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

func (s *ReportsService) downloadGeneratedReport(ctx context.Context, reportURL string, format models.ReportFormat) (*models.GeneratedReport, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reportURL, nil)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	reportBody, err := readGeneratedReportBody(bytes.NewReader(bodyBytes), format)
	if err != nil {
		return nil, err
	}

	return &models.GeneratedReport{
		URL:  reportURL,
		Body: json.RawMessage(reportBody),
	}, nil
}

func (s *ReportsService) doRequest(req *http.Request) ([]byte, error) {
	res, err := s.client.httpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode/100 != 2 {
		errBody, _ := io.ReadAll(res.Body)
		return nil, newAPIError(res.Status, res.StatusCode, string(errBody))
	}

	return io.ReadAll(res.Body)
}

func readGeneratedReportBody(body io.Reader, format models.ReportFormat) ([]byte, error) {
	if format != models.ReportFormatGZIPJSON {
		return io.ReadAll(body)
	}

	gzipReader, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = gzipReader.Close()
	}()

	return io.ReadAll(gzipReader)
}

func (r *Report) ReportID() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.details.ReportID
}

func (r *Report) Details() models.ReportDetails {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.details
}

func (r *Report) Refresh(ctx context.Context) (*models.ReportDetails, error) {
	reportID := r.ReportID()
	details, err := r.service.getReportDetails(ctx, r.profileID, reportID)
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	r.details = *details
	refreshed := r.details
	r.mu.Unlock()

	return &refreshed, nil
}

func (r *Report) RequestState(ctx context.Context) (models.ReportStatus, error) {
	details, err := r.Refresh(ctx)
	if err != nil {
		return "", err
	}

	return details.Status, nil
}

func (r *Report) IsTerminal(ctx context.Context) (bool, error) {
	state, err := r.RequestState(ctx)
	if err != nil {
		return false, err
	}

	return state == models.ReportStatusCompleted || state == models.ReportStatusFailed, nil
}

func (r *Report) GeneratedReport(ctx context.Context) (*models.GeneratedReport, error) {
	details := r.Details()
	if !details.IsCompleted() || !details.HasDownloadURL() {
		refreshed, err := r.Refresh(ctx)
		if err != nil {
			return nil, err
		}
		details = *refreshed
	}

	if !details.IsCompleted() {
		if details.FailureReason != nil && *details.FailureReason != "" {
			return nil, fmt.Errorf("report is not completed: %s (%s)", details.Status, *details.FailureReason)
		}
		return nil, fmt.Errorf("report is not completed: %s", details.Status)
	}

	if !details.HasDownloadURL() {
		return nil, errors.New("completed report does not include a download URL")
	}

	return r.service.downloadGeneratedReport(ctx, *details.URL, details.Configuration.Format)
}
