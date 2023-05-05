package main

type PaymentCreditCard struct {
}

func (pcc *PaymentCreditCard) Payment(amount int) string {

	// TODO: DO SOME THING

	resultMessage := "신용카드로 결제하였습니다."
	return resultMessage
}
