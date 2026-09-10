package main

import (
	"fmt"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Role     string
	IsActive bool
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))

	for i, item := range items {
		result[i] = transform(item)
	}

	return result
}

func main() {
	users := []User{
		{1, "Alice", "admin", true},
		{2, "Bob", "user", false},
		{3, "Charlie", "user", true},
		{4, "Diana", "admin", true},
	}

	isActive := func(u User) bool { return u.IsActive }
	isAdmin := func(u User) bool { return u.Role == "admin" }

	activeAdmins := Filter(Filter(users, isActive), isAdmin)
	fmt.Printf("active admins: %v\n", activeAdmins)

	names := Map(activeAdmins, func(u User) string {
		return strings.ToUpper(u.Name)
	})

	fmt.Printf("active admin names: %v\n", names)
}
