package main

import "fmt"

func main() {
	var payment Payment

	// 비트코인 결제
	payment = new(PaymentBitcoin)
	result := payment.Payment(10000)
	fmt.Println(result)

	// 페이팔 결제
	payment = new(PaymentPaypal)
	result = payment.Payment(200000000)
	fmt.Println(result)
}
