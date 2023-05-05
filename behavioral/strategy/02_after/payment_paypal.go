package main

type PaymentPaypal struct {
}

func (pp *PaymentPaypal) Payment(amount int) string {

	// TODO: DO SOME THING

	resultMessage := "페이팔로 결제하였습니다."
	return resultMessage
}
