package main

import "fmt"

func grantBenefits(tier string) {
	if tier != "GOLD" && tier != "SILVER" && tier != "BRONZE" {
		fmt.Println("unknown tier: no benefits")
		return
	}

	fmt.Printf("%s member gets:\n", tier)

	switch tier {
	case "GOLD":
		fmt.Println("- Priority support queue")
		fallthrough
	case "SILVER":
		fmt.Println("- Free monthly report")
		fallthrough
	case "BRONZE":
		fmt.Println("- 5% cashback")
	}
}

func main() {
	grantBenefits("GOLD")
	fmt.Println()
	grantBenefits("SILVER")
	fmt.Println()
	grantBenefits("BRONZE")
	fmt.Println()
	grantBenefits("FREE")
}
