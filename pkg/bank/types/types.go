package types

//Money представляет собой денежную сумму в минимальной денежных еденицах (центы, копейки, дирамы и т.д)
type Money int

//Currency представляет код валюты
type Currency string

// Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string


//Card представляет информацию о платёжной карте
type Card struct {
	Type    string
	ID		int
	PAN		PAN
	Number string
	Balance Money
	MinBalance Money
	Currency Currency
	Color	string
	Name	string
	Active 	bool
}

//Payment представляет информацию о платеже
type Payment struct {
	ID int
	Amount Money 
}

//PaymentSource представляет структуру кода платежа и вывода карт платежа
type PaymentSource struct{
	Type string
	Number string
	Balance Money
}