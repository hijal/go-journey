package main

import "fmt"

func HasRole(userRole string, allowedRoles ...string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasRole("editor", "admin", "moderator", "editor"))
	fmt.Println(HasRole("guest", "admin", "moderator"))
}
