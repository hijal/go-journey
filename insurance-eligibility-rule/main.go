package main

import "fmt"

type Claim struct {
	Month  string
	Amount int
}

type Applicant struct {
	Name       string
	Age        int
	HasLicense bool
	Claims     []Claim
}

func main() {
	applicant := Applicant{
		Name:       "John Doe",
		Age:        25,
		HasLicense: true,
		Claims: []Claim{{
			Month:  "March",
			Amount: 40000,
		}},
	}

	hasBigClaim := len(applicant.Claims) > 0 && applicant.Claims[0].Amount >= 100000
	isAdult := applicant.Age >= 18
	isEligible := isAdult && applicant.HasLicense && !hasBigClaim

	fmt.Printf("Applicant:  %s\n", applicant.Name)
	fmt.Printf("Adult:      %v\n", isAdult)
	fmt.Printf("Big claim:  %v\n", hasBigClaim)
	fmt.Printf("Eligible:   %v\n", isEligible)

	if !isEligible {
		fmt.Println("Decision: referred to a human underwriter")
		return
	}
	fmt.Println("Decision: auto-approved, quote 9000 BDT")
}
