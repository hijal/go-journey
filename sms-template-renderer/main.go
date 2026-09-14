package main

import (
	"fmt"
	"strings"
)

const orderTemplate = `নমস্কার {{name}},
আপনার অর্ডার {{order_id}} গ্রহণ করা হয়েছে (মোট {{total}} টাকা)।
ট্র্যাকিং লিংক: {{tracking_url}}`

func renderOrderSMS(name, orderID, total, trackingURL string) string {
	r := strings.NewReplacer(
		"{{name}}", name,
		"{{order_id}}", orderID,
		"{{total}}", total,
		"{{tracking_url}}", trackingURL,
	)
	return r.Replace(orderTemplate)
}

func main() {
	fmt.Println(renderOrderSMS("রাফি", "ORD-9012", "1450", "https://example.com/t/9012"))
	fmt.Println("---")
	fmt.Println(renderOrderSMS("মিতু", "ORD-9013", "899", "https://example.com/t/9013"))
}
