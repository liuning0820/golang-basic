package main

import (
	"fmt"
	"os/user"
	"time"
)

type Item struct {
    ID string
}

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    user.User
    Items        []Item
    CurrencyCode string
    isLocked     bool
}

func (c *Cart) TotalPrice() (string, error) {
    // ...
    return "12",nil
}

func main() {

	newCart:= Cart{}

	totalPrice, err :=newCart.TotalPrice()

	if err != nil {
        fmt.Printf("impossible to compute price of the cart: %s", err)
        return
    }
    fmt.Println("Total Price", totalPrice)

}