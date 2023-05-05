package main

type Payment interface {
	Payment(amount int) string
}

type PaymentMethod string

var (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodPaypal     PaymentMethod = "paypal"
	PaymentMethodBitcoin    PaymentMethod = "bitcoin"
)
