package main

import (
	"errors"
	"fmt"
)

type OrderStore interface {
	Status(id string) (string, error)
	SaveStatus(id, status string) error
}

var ErrNotFound = errors.New("order not found")

type fakeStore struct {
	db map[string]string
}

func (f *fakeStore) Status(id string) (string, error) {
	status, ok := f.db[id]
	if !ok {
		return "", fmt.Errorf("lookup %s: %w", id, ErrNotFound)
	}
	return status, nil
}

func (f *fakeStore) SaveStatus(id, status string) error {
	f.db[id] = status
	return nil
}

func notifyFailedOrders(store OrderStore, ids []string) {
	for _, id := range ids {
		status, err := store.Status(id)

		if err != nil {
			if errors.Is(err, ErrNotFound) {
				fmt.Println("skip", id, "(not found)")
				continue
			}
			fmt.Println("skip", id, "error:", err)
			continue
		}

		if status == "FAILED" {
			if err := store.SaveStatus(id, "NOTIFIED"); err != nil {
				fmt.Println("save failed for", id, ":", err)
				continue
			}
			fmt.Println("notified customer for", id)
		} else {
			fmt.Println("no action for", id, "status:", status)
		}
	}
}

func main() {
	store := &fakeStore{db: map[string]string{
		"ORD-1": "FAILED",
		"ORD-2": "SETTLED",
	}}

	notifyFailedOrders(store, []string{"ORD-1", "ORD-2", "ORD-9"})
}
