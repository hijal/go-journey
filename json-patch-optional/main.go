package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

type UserUpdateRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Age   *int    `json:"age,omitempty"`
}

func applyUpdate(existing *User, req UserUpdateRequest) {
	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.Email != nil {
		existing.Email = *req.Email
	}

	if req.Age != nil {
		existing.Age = *req.Age
	}
}

func main() {
	user := &User{Name: "bob", Email: "bob@example.com", Age: 30}

	payload := `{"age": 31}`

	var req UserUpdateRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}

	applyUpdate(user, req)
	fmt.Printf("%+v\n", user)
}
