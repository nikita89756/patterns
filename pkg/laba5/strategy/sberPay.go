package payment

import "fmt"

type SberPay struct{}

func (s SberPay) writeOffMoney(money float64)string{
	return fmt.Sprintf("Списание с карты сбербанка на сумму %.2f",money)
}