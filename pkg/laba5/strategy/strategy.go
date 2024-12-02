package payment

type PaymentStrategy interface{
	writeOffMoney(money float64) string
}