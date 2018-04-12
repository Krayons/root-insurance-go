package insurance

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	applicaitonstURL = "applications"
)

type ApplicationService struct {
	*Client
}

type CreateApplicationRequest struct {
	MonthlyPremium int    `json:"monthly_premium"`
	QuotePackageID string `json:"quote_package_id"`
	SerialNumber   string `json:"serial_number"`
	PolicyholderID string `json:"policyholder_id"`
}

type ApplicationResponse struct {
	ApplicationID  string `json:"application_id"`
	PolicyholderID string `json:"policyholder_id"`
	PackageName    string `json:"package_name"`
	SumAssured     int    `json:"sum_assured"`
	MonthlyPremium int    `json:"monthly_premium"`
	BasePremium    int    `json:"base_premium"`
	Module         struct {
		Type         string `json:"type"`
		Make         string `json:"make"`
		Model        string `json:"model"`
		SerialNumber string `json:"serial_number"`
	} `json:"module"`
	CreatedAt time.Time `json:"created_at"`
}

func (ps *ApplicationService) CreateApplication(request CreateApplicationRequest) (*ApplicationResponse, error) {
	jsonReq, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, ps.rootURL+applicaitonstURL, bytes.NewBuffer(jsonReq))
	resp, err := ps.Do(req)
	if err != nil {
		return nil, err
	}
	results := ApplicationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}
