package Structs

import "fmt"

type Letter struct {
	senderAddress    string
	recipientAddress string
}

func (l Letter) GetSenderAddress() string {
	return l.senderAddress
}
func (l Letter) GetRecipientAddress() string {
	return l.recipientAddress
}
func (l Letter) Send() {
	fmt.Printf("Доставка листа від %s до %s\n", l.senderAddress, l.recipientAddress)
}
