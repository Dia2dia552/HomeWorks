package Structs

type Parcel interface {
	GetSenderAddress() string
	GetRecipientAddress() string
	Send()
}
