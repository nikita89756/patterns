package payment

import "fmt"

type TBankPay struct{}

func (t TBankPay) writeOffMoney(money float64)string{
	return fmt.Sprintf("Списание с карты T-Bank на сумму %.2f",money)
}