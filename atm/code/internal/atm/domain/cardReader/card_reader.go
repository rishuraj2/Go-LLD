package cardreader

import "errors"

type CardReader struct {
	cardNumber string
}

func NewCardReader() *CardReader {
	return &CardReader{}
}

func (this *CardReader) ReadCard(number string) error {
	if len(number) != 10 {
		return errors.New("invalid card")
	}

	this.cardNumber = number
	return nil
}

func (this *CardReader) GetCardNumber() string {
	return this.cardNumber
}

func (this *CardReader) EjectCard() {
	this.cardNumber = ""
}
