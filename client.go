package insurance

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	start := time.Now()

	records := GadgetModels{}
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	return records, nil
}

// Create a new root client from a api key
func Create(apiKey string) (*Client, error) {
	client := Client{}
	client.apiKey = apiKey
	client.rootURL = "https://sandbox.root.co.za/v1/insurance/"
	client.Client = &http.Client{}
	client.Gadget = Gadget{&client}
	client.QuoteService = &QuoteService{&client}
	return &client, nil
}

// Inject api headers into all requests
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("X-Client-Header", "Go root client")
	req.Header.Set("Content-Type", "application/json")
	return client.Client.Do(req)
}
