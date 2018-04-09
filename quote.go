package insurance

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	quoteURL = "quotes"
)

type QuoteService struct {
	*Client
}

type QuoteType struct {
	Type string `json:"type"`
}

type GadgetQuote struct {
	QuotePackageID   string `json:"quote_package_id"`
	PackageName      string `json:"package_name"`
	SumAssured       int    `json:"sum_assured"`
	BasePremium      int    `json:"base_premium"`
	SuggestedPremium int    `json:"suggested_premium"`
	Module           struct {
		Type  string `json:"type"`
		Make  string `json:"make"`
		Model string `json:"model"`
	} `json:"module"`
	CreatedAt time.Time `json:"created_at"`
}

type FuneralQuote struct {
	QuotePackageID   string `json:"quote_package_id"`
	PackageName      string `json:"package_name"`
	SumAssured       int    `json:"sum_assured"`
	BasePremium      int    `json:"base_premium"`
	SuggestedPremium int    `json:"suggested_premium"`
	Module           struct {
		Type               string `json:"type"`
		HasSpouse          bool   `json:"has_spouse"`
		NumberOfChildren   int    `json:"number_of_children"`
		ExtendedFamilyAges []int  `json:"extended_family_ages"`
	} `json:"module"`
	CreatedAt time.Time `json:"created_at"`
}

// GadgetQuoteRequest used for when query gadet quotes
type GadgetQuoteRequest struct {
	QuoteType
	ModelName string `json:"model_name"`
}

type FuneralQuoteRequest struct {
	Type               string `json:"type"`
	CoverAmount        int    `json:"cover_amount"`
	HasSpouse          bool   `json:"has_spouse"`
	NumberOfChildren   int    `json:"number_of_children"`
	ExtendedFamilyAges []int  `json:"extended_family_ages"`
}

func (qs *QuoteService) GetGadgetQuote(request GadgetQuoteRequest) ([]GadgetQuote, error) {
	jsonReq, _ := json.Marshal(request)
	req, _ := http.NewRequest(http.MethodPost, qs.rootURL+quoteURL, bytes.NewBuffer(jsonReq))
	resp, err := qs.Do(req)
	if err != nil {
		return nil, err
	}
	results := []GadgetQuote{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}
