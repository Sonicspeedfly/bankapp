package stats

import (
	"github.com/Sonicspeedfly/bankapp/pkg/bank/types"
	"fmt"
)

func ExamleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
		},
		{
			Amount: 15_000,
		},
	}
	sred := Avg(payments)
	fmt.Println(sred)

	//Output:
	//17500
}