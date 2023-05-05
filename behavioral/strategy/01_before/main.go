package main

import "fmt"

func main() {
	result := Payment(129000, PaymentMethodBitcoin)
	fmt.Println(result)

	result = Payment(129000, PaymentMethodPaypal)
	fmt.Println(result)
}
