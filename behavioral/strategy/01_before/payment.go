package main

type PaymentMethod string

var (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodPaypal     PaymentMethod = "paypal"
	PaymentMethodBitcoin    PaymentMethod = "bitcoin"
)

func Payment(amount int, paymentMethod PaymentMethod) string {
	var resultMessage string

	if paymentMethod == PaymentMethodCreditCard {
		// TODO: DO SOMETHING...
		resultMessage = "신용카드로 결제하였습니다."
	} else if paymentMethod == PaymentMethodPaypal {
		// TODO: DO SOME THING
		resultMessage = "페이팔로 결제하였습니다."
	} else {
		// TODO: DO SOME THING
		resultMessage = "비트코인으로 결제하였습니다."
	}

	return resultMessage
}
