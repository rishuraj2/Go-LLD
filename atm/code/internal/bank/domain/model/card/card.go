package card

type Card struct {
	cardNumber    string
	accountNumber string
	pin           string
}

func NewCard(cardNumber, accountNumber, pin string) Card {
	return Card{
		cardNumber:    cardNumber,
		accountNumber: accountNumber,
		pin:           pin,
	}
}

func (this *Card) GetCardNumber() string {
	return this.cardNumber
}

func (this *Card) GetAccountNumber() string {
	return this.accountNumber
}

func (this *Card) GetPin() string {
	return this.pin
}
