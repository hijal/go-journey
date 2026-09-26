package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type productPage struct {
	stock   int
	price   int
	reviews float64
}

var errReviewsDown = errors.New("reviews service unavailable")

func fetchStock(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 34, nil
}

func fetchPrice(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 2450, nil
}

func fetchReviews(sku string) (float64, error) {
	time.Sleep(60 * time.Millisecond)
	return 0, errReviewsDown
}

func main() {
	const sku = "SKU-77120"

	start := time.Now()

	var (
		page productPage
		wg   sync.WaitGroup
		errs [3]error
	)

	wg.Go(func() {
		v, err := fetchStock(sku)

		if err != nil {
			errs[0] = fmt.Errorf("stock: %w", err)
			return
		}
		page.stock = v
	})

	wg.Go(func() {
		v, err := fetchPrice(sku)

		if err != nil {
			errs[1] = fmt.Errorf("price: %w", err)
			return
		}

		page.price = v
	})

	wg.Go(func() {
		v, err := fetchReviews(sku)
		if err != nil {
			errs[2] = fmt.Errorf("reviews: %w", err)
			return
		}
		page.reviews = v
	})

	wg.Wait()

	fmt.Printf("page %s: stock=%d price=%d reviews=%.1f\n",
		sku, page.stock, page.price, page.reviews)
	fmt.Println("concurrent fetch took",
		time.Since(start).Round(10*time.Millisecond), "(sequential would be ~270ms)")

	if err := errors.Join(errs[:]...); err != nil {
		fmt.Println("degraded:", err)
		fmt.Println("reviews down?", errors.Is(err, errReviewsDown))
	}
}
