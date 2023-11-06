package unraveling_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestCreateSimpleOrder(t *testing.T) {
	err := createSimpleOrder()
	if err != nil {
		fmt.Println("failed to create order")
	}
}

func TestCreateOrder(t *testing.T) {
	err := createOrder()
	if err != nil {
		fmt.Println("failed to create order")
	}
}

// Case 1
func createSimpleOrder() error {
	var err error

	err = adjustItemStock()
	if err != nil {
		return err
	}

	err = adjustPromoQuota()
	if err != nil {
		return err
	}

	fmt.Println("order created successfully")
	return nil
}

// Concurrency Implementation
func createOrder() error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err1 := adjustItemStock()
		if err == nil {
			err = err1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err2 := adjustPromoQuota()
		if err == nil {
			err = err2
		}
	}()

	wg.Wait()

	if err != nil {
		return err
	}

	fmt.Println("order created successfully")
	return nil
}

func adjustItemStock() error {
	return nil
}

func adjustPromoQuota() error {
	return nil
}
