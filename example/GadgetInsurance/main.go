package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Krayons/root-insurance-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	rootInsurance, err := insurance.Create("apiKey") // api key needs to be replaced
	if err != nil {
		panic(err)
	}
	gadgets, err := rootInsurance.Gadget.Gadgets()
	if err != nil {
		panic(err)
	}

	// get a random gadget from all the insurable gadgets
	randomIdx := rand.Intn(len(gadgets))
	gadget := gadgets[randomIdx]
	fmt.Printf("We are going to be randomly insuring a %s\n", gadget.Name)
	fmt.Println("Retriving a quote")
	qr := rootInsurance.QuoteService.CreateQuoteRequest(gadget.Name)
	quotes, err := rootInsurance.QuoteService.GetGadgetQuote(qr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("We recieved %d quotes for the %s\n", len(quotes), gadget.Name)
	for idx, quote := range quotes {
		fmt.Printf("Quote %d is %s\n", idx+1, quote.PackageName)
		fmt.Printf("With a Suggested Premium of R%v/pm for the value of R%v\n", quote.SuggestedPremium/100, quote.SumAssured/100)
	}
}
