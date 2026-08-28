package atm

import (
	cardreader "atm/internal/atm/domain/cardReader"
	cashdispenser "atm/internal/atm/domain/cashDispenser"
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
	"atm/internal/bank/app/service"
	"sync"
)

type ATMState interface {
	InsertCard(ctx *ATM, cardNumber string) error
	Authenticate(ctx *ATM, pin string) error
	Transact(ctx *ATM, transactionType enum.TransactionType) error
	Process(ctx *ATM) error
	Dispense(ctx *ATM) error
	EjectCard(ctx *ATM) error
}

type ATM struct {
	state           ATMState
	dispenser       *cashdispenser.CashDispenser
	reservoir       *cashreservoir.CashReservoir
	cardReader      *cardreader.CardReader
	server          *service.BankService
	accountNumber   string
	transactionType enum.TransactionType
	transactionID   string
}

var (
	instance *ATM
	once     sync.Once
)

func NewATM(dispenser *cashdispenser.CashDispenser, reservoir *cashreservoir.CashReservoir, cardReader *cardreader.CardReader, server *service.BankService) *ATM {
	once.Do(func() {
		instance = &ATM{
			state:      nil,
			dispenser:  dispenser,
			reservoir:  reservoir,
			cardReader: cardReader,
			server:     server,
		}
	})

	return instance
}

func (this *ATM) GetCardReader() *cardreader.CardReader {
	return this.cardReader
}

func (this *ATM) GetBankServer() *service.BankService {
	return this.server
}

func (this *ATM) GetAccountNumber() string {
	return this.accountNumber
}

func (this *ATM) GetTransactionID() string {
	return this.transactionID
}

func (this *ATM) GetTransactionType() enum.TransactionType {
	return this.transactionType
}

func (this *ATM) SetState(state ATMState) {
	this.state = state
}

func (this *ATM) SetAccountNumber(number string) {
	this.accountNumber = number
}

func (this *ATM) SetTransactionID(id string) {
	this.transactionID = id
}

func (this *ATM) SetTransactionType(trxnType enum.TransactionType) {
	this.transactionType = trxnType
}
