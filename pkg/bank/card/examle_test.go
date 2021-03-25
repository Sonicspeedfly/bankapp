package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
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
	maxes := PaymentSources(payments)
	for _, max := range maxes {
		fmt.Println(max.Number)
	}
	//Output:
	//15
	//18
}