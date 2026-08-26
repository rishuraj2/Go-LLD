package service

import (
	"atm/internal/bank/app/usecase"
	"atm/internal/bank/domain/enum"
	"atm/internal/bank/domain/model/transaction"
	"atm/internal/bank/repository"
)

type BankService struct {
	createAccountUseCase      *usecase.CreateAccountUseCase
	createCardUseCase         *usecase.CreateCardUseCase
	authenticateCardUseCase   *usecase.AuthenticateCardUseCase
	creditAmountUseCase       *usecase.CreditAmountUseCase
	debitAmountUseCase        *usecase.DebitAmountUseCase
	transactionHistoryUseCase *usecase.TransactionHistoryUseCase
	transactionUseCase        *usecase.TransactionUseCase
}

func NewBankService(repo repository.BankRepository) *BankService {
	return &BankService{
		createAccountUseCase:      usecase.NewCreateAccountUseCase(repo),
		createCardUseCase:         usecase.NewCreateCardUseCase(repo),
		authenticateCardUseCase:   usecase.NewAuthenticateCardUseCase(repo),
		creditAmountUseCase:       usecase.NewCreditAmountUseCase(repo),
		debitAmountUseCase:        usecase.NewDebitAmountUseCase(repo),
		transactionHistoryUseCase: usecase.NewTransactionHistoryUseCase(repo),
		transactionUseCase:        usecase.NewTransactionUseCase(repo),
	}
}

func (this *BankService) CreateAccount(accountType enum.AccountType) string {
	return this.createAccountUseCase.Execute(accountType)
}

func (this *BankService) CreateCard(accountNumber, pin string) (string, error) {
	return this.createCardUseCase.Execute(accountNumber, pin)
}

func (this *BankService) AuthenticateCard(cardNumber, pin string) (string, error) {
	return this.authenticateCardUseCase.Execute(cardNumber, pin)
}

func (this *BankService) CreditAmount(accountNumber string, amount float64) (string, error) {
	return this.creditAmountUseCase.Execute(accountNumber, amount)
}

func (this *BankService) DebitAmount(accountNumber string, amount float64) (string, error) {
	return this.debitAmountUseCase.Execute(accountNumber, amount)
}

func (this *BankService) GetTransactionHistory(accountNumber string) []transaction.Transaction {
	return this.transactionHistoryUseCase.Execute(accountNumber)
}

func (this *BankService) GetTransactionByID(transactionID string) (transaction.Transaction, error) {
	return this.transactionUseCase.Execute(transactionID)
}
