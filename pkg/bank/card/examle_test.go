package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	payments := []types.Card{
		{
			PAN: "15", 
			Active: true,
			Balance: 15_000,
		},
		{
			PAN: "99",
			Active: false,
			Balance: 30_000,
		},
		{
			PAN: "33",
			Active: true,
			Balance: -15_000,
		},
	}
	maxes := PaymentSources(payments)
	for _, max := range maxes {
		fmt.Println(max.Number)
	}
	//Output:
	//15
}