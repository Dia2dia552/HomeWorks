package Structs

import "fmt"

type Box struct {
	senderAddress    string
	recipientAddress string
}

func (b Box) GetSenderAddress() string {
	return b.senderAddress
}
func (b Box) GetRecipientAddress() string {
	return b.recipientAddress
}
func (b Box) Send() {
	fmt.Printf("Доставка коробки від %s до %s\n", b.senderAddress, b.recipientAddress)
}
