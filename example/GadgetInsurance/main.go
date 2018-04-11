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
	rootInsurance, err := insurance.Create("sandbox_ZjFhZTI3M2UtMmUzZS00ZjFkLTg1YjItNGE2M2ViMzc2MDdhLnhsaHExNTlpdkdPQmVYT3gtZVFLb2dwTmxvZk9MZ0xG") // api key needs to be replaced
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
	fmt.Println("Retrieving all policyholders")
	policyholders, _ := rootInsurance.PolicyholderService.GetAllPolicyholders()
	noHolders := len(policyholders)
	fmt.Printf("Found %v policyholders\n", noHolders)
	var policyholder insurance.Policyholder
	if noHolders < 0 {
		fmt.Println("Creating a policyholder")
		policyholder = insurance.Policyholder{}
	} else {
		fmt.Println("Selecting a random policy holder")
		policyholder = policyholders[rand.Intn(noHolders)]
	}
	fmt.Printf("Policyholder chosen: %s %s", policyholder.FirstName, policyholder.LastName)

}
