package unraveling_goroutine

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCreateSimpleNotifiedOrder(t *testing.T) {
	err := createSimpleNotifiedOrder()
	if err != nil {
		fmt.Println("failed to create order")
	}
}

func TestCreateNotifiedOrder(t *testing.T) {
	err := createNotifiedOrder()
	if err != nil {
		fmt.Println("failed to create order")
	}
}

// Case 2
func createSimpleNotifiedOrder() error {
	var err error

	err = adjustItemStock()
	if err != nil {
		return err
	}

	err = adjustPromoQuota()
	if err != nil {
		return err
	}

	err = notifyCustomerNewOrder()
	if err != nil {
		fmt.Println("failed to send new order notification to customer")
	}

	fmt.Println("order created successfully")
	return nil
}

// Asynchronous Implementation
func createNotifiedOrder() error {
	var err error

	err = adjustItemStock()
	if err != nil {
		return err
	}

	err = adjustPromoQuota()
	if err != nil {
		return err
	}

	go func() {
		err = notifyCustomerNewOrder()
		if err != nil {
			fmt.Println("failed to send new order notification to customer")
		}
	}()

	fmt.Println("order created successfully")
	time.Sleep(1 * time.Second)
	return nil
}

func notifyCustomerNewOrder() error {
	return errors.New("exception")
}
