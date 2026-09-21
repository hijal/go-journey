package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

type Order struct {
	ID         string
	CustomerID string
	Email      string
	TotalCents int64
	Status     string
}

type OrderRepository interface {
	Save(ctx context.Context, o Order) error
}

type Notifier interface {
	Notify(ctx context.Context, to, message string) error
}

type OrderService struct {
	repo     OrderRepository
	notifier Notifier
}

func NewOrderService(repo OrderRepository, notifier Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

var ErrEmptyOrder = errors.New("order total must be positive")

func (s *OrderService) PlaceOrder(ctx context.Context, o Order) error {
	if o.TotalCents <= 0 {
		return ErrEmptyOrder
	}

	o.Status = "placed"
	if err := s.repo.Save(ctx, o); err != nil {
		return fmt.Errorf("save order %s: %w", o.ID, err)
	}
	msg := fmt.Sprintf("Your order %s has been placed.", o.ID)

	if err := s.notifier.Notify(ctx, o.Email, msg); err != nil {
		log.Printf("notify %s: %v", o.Email, err)
	}
	return nil
}

type MemoryOrderRepo struct {
	mu     sync.Mutex
	orders map[string]Order
}

func NewMemoryOrderRepo() *MemoryOrderRepo {
	return &MemoryOrderRepo{orders: make(map[string]Order)}
}

func (r *MemoryOrderRepo) Save(_ context.Context, o Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[o.ID]; exists {
		return fmt.Errorf("order %s already exist", o.ID)
	}
	r.orders[o.ID] = o
	return nil
}

type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(_ context.Context, to, message string) error {
	fmt.Printf("Email to=%s: %s\n", to, message)
	return nil
}

func main() {
	svc := NewOrderService(NewMemoryOrderRepo(), ConsoleNotifier{})
	ctx := context.Background()

	order := Order{ID: "ord-1001", CustomerID: "c-7", Email: "example@example.com", TotalCents: 1466}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		log.Fatal(err)
	}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		fmt.Println("error:", err)
	}
}
