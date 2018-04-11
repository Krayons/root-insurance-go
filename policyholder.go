package insurance

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	endPointURL = "policyholders"
)

type PolicyholderService struct {
	*Client
}

type Policyholder struct {
	PolicyholderID string `json:"policyholder_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ID             struct {
		Type    string `json:"type"`
		Number  string `json:"number"`
		Country string `json:"country"`
	} `json:"id"`
	DateOfBirth string      `json:"date_of_birth"`
	Gender      string      `json:"gender"`
	Email       string      `json:"email"`
	Cellphone   interface{} `json:"cellphone"`
	AppData     struct {
		Company string `json:"company"`
	} `json:"app_data"`
	PolicyIds []interface{} `json:"policy_ids"`
	CreatedAt time.Time     `json:"created_at"`
	CreatedBy struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"created_by"`
}

func (ps *PolicyholderService) GetAllPolicyholders() ([]Policyholder, error) {
	req, _ := http.NewRequest(http.MethodGet, ps.rootURL+endPointURL, nil)
	resp, err := ps.Do(req)
	if err != nil {
		return nil, err
	}
	results := []Policyholder{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}
