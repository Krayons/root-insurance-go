package insurance

import (
	"encoding/json"
	"net/http"
)

const (
	gadgets = "modules/root_gadgets/models"
)

// Move everything below here to some where else
type GadgetModels []struct {
	Make  string `json:"make"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Gadget struct {
	*Client
}

// Client for root insurance interface
type Client struct {
	apiKey string
	Gadget
	*QuoteService
	*http.Client
	rootURL string
}

// Gadgets returns current gadgets root insures
func (client *Gadget) Gadgets() (GadgetModels, error) {
	req, _ := http.NewRequest(http.MethodGet, client.rootURL+gadgets, nil)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	records := GadgetModels{}
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, err
	}
	return records, nil
}

// Create a new root client from a api key
func Create(apiKey string, options ...func(*Client) error) (*Client, error) {
	client := Client{}
	client.apiKey = apiKey
	client.rootURL = "https://sandbox.root.co.za/v1/insurance/"
	client.Client = &http.Client{}
	client.Gadget = Gadget{&client}
	client.QuoteService = &QuoteService{&client}
	for _, opt := range options {
		err := opt(&client)
		if err != nil {
			return nil, err
		}
	}
	return &client, nil
}

// Inject api headers into all requests
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("X-Client-Header", "Go root client")
	req.Header.Set("Content-Type", "application/json")
	return client.Client.Do(req)
}
