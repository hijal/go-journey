package main

import "fmt"

type Address struct {
	City    string
	Street  string
	ZipCode string
}

type Customer struct {
	Name        string
	Home        Address
	Office      Address
	PhoneNumber []string
}

func main() {
	c := Customer{
		Name: "John Doe",
		Home: Address{
			City:    "New York",
			Street:  "123 Main St",
			ZipCode: "10001",
		},
		Office: Address{
			City:    "New York",
			Street:  "456 Office Rd",
			ZipCode: "10002",
		},
		PhoneNumber: []string{"123-456-7890", "987-654-3210"},
	}

	fmt.Println("Home city:", c.Home.City)
	fmt.Println("Office city:", c.Office.City)
	c.Home.City = "Auckland"
	fmt.Printf("Updated home: %+v\n", c.Home)
}
