package main

import "fmt"

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.92,
	"BDT": 117.50,
	"GBP": 0.79,
}

func convert(amount float64, from, to string) (float64, error) {
	fromRate, ok := rates[from]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", from)
	}

	toRate, ok := rates[to]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", to)
	}

	return amount / fromRate * toRate, nil
}

func applySpread(r map[string]float64, pct float64) {
	for code := range r {
		r[code] *= 1 + pct/100
	}
}

func main() {
	bdt, err := convert(100, "USD", "BDT")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("100 USD = %.2f BDT\n", bdt)

	if _, err := convert(50, "USD", "XAU"); err != nil {
		fmt.Println("error:", err)
	}

	applySpread(rates, 2)

	fmt.Printf("After 2%% spread, USD->BDT rate is %.2f\n", rates["BDT"])
}
