package main

type PaymentBitcoin struct {
}

func (pb *PaymentBitcoin) Payment(amount int) string {

	// TODO: DO SOME THING

	resultMessage := "비트코인으로 결제하였습니다."
	return resultMessage
}
