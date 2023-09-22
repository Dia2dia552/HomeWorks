package Structs

func Delivery() {
	box := Box{senderAddress: "Lviv", recipientAddress: "Odessa"}
	letter := Letter{senderAddress: "Kyiv", recipientAddress: "Kherson"}

	sortingDepartment := func(p Parcel) {
		p.Send()
	}
	sortingDepartment(box)
	sortingDepartment(letter)

}
