package card

import "bank/pkg/bank/types"

//Max возвращает максимальный платёж
func Max(payments []types.Payment) types.Payment{
	max:=payments[0]
	for _, payment := range payments {
		if (max.Amount < payment.Amount){
			max=payment
		}
	}
	return max
}

//PaymentSources выводит карты платежа
func PaymentSources(cards []types.Card) []types.PaymentSource{
	var operations []types.PaymentSource
	for _, card := range cards {
		if ((card.Balance > 0)&&(card.Active == true)){
			operations = append(operations, types.PaymentSource{Type: "card", Number: card.Number, Balance: card.Balance})
			}
		}
		return operations	
	}