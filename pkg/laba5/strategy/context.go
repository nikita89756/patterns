package payment

import (
	"fmt"
	Mlogger "misis/pkg/laba3/singleton"
)

type PaymentCard struct{
	cardNumber int
	cardDate int
	cvv int
	PaymentMethod PaymentStrategy
	logger Mlogger.Mlogger
}

func NewPaymentCard(cardNumber,cardDate,cvv int , bank string,logger Mlogger.Mlogger) *PaymentCard{
	var paymentMethod PaymentStrategy
	switch bank{
	case "sber":
		paymentMethod = SberPay{}
	case "tbank":
		paymentMethod = TBankPay{}
	}
	return &PaymentCard{cardNumber: cardNumber,cardDate: cardDate,cvv: cvv,PaymentMethod: paymentMethod,logger: logger}

}
func (p *PaymentCard)Pay(money float64){
	fmt.Println(p.logger.LogInfo(p.PaymentMethod.writeOffMoney(money)))
}