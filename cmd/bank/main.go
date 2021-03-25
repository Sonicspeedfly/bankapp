package main

import (
	"fmt"
	"bank/pkg/bank/types"
	"bank/pkg/bank/card"
   )
 
func main() {
	var NewCard types.PaymentSource
	payments := []types.Card{
		{
			Number: "15", 
			Active: true,
			Balance: 15_000,
		},
		{
			Number:	"99",
			Active: false,
			Balance: 30_000,
		},
		{
			Number: "18",
			Active: true,
			Balance: 15_000,
		},
	}
	maxes := card.PaymentSources(payments)
	for _, max := range maxes {
		NewCard.Number=max.Number
		fmt.Println(NewCard.Number)
	}
}