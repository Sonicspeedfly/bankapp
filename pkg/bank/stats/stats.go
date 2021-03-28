package stats

import (
	"github.com/Sonicspeedfly/bankapp/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа .
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		sum += payment.Amount
	}
	sred := sum / 2
	return sred
}