# Functional Requirements
- ATM supports three types of transactions: cash withdrawl, cash deposite and balance enquiry
- User authentication should be card-based with pin
- There is a bank service that authenticates and authorizes all the actions
- Dispense cash in largest denomination first (₹2000, ₹500, ₹200, ₹100, ₹50, ₹20, ₹10)
- Validate both account balance and ATM cash inventory before dispensing
- Track ATM state transition (ideal, card inserted, authenticated, transacting, processing, dispensing, Eject)
- Before making API, hardcode the sequence in main function.

# Non Functional Requirements
- The design should follow object-oriented principles with clear seperation of concerns.
- System should be modular and extensible to support new transaction types and denomination.
- Code should be thread-safe for concurrent access.
- Components should be testable in isolation.
- Handle error gracefully with appropriate error message.
- Financial operation should always follow validation-before-commit: always verify dispensing before debiting.

# Core Entities
- AccountType (Enum) (CURRENT, SAVINGS)
- ATMState (Enum) (IDEAL, CARD_INSERTED, AUTHENTICATED, TRANSACTING, PROCESSING, DISPENSING, CARD_EJECTED)
- TransactionType (Enum) (WITHDRAWL, DEPOSIT, BALANCE_INQUIRY)
- CashDenomination (Enum) (TWO_THOUSAND, FIVE_HUNDRED, TWO_HUNDRED, ONE_HUNDRED, FIFTY, TWENTY, TEN)
- Card
- Account
- Transaction
- CashDispenser
- BankServer
- ATM

# Responsibility
- **AccountType**: Enum having "CURRENT" and "SAVINGS" as value. Used by **Account**
- **ATMState**: Enum having "IDEAL", "CARD_INSERTED", "AUTHENTICATED", "TRANSACTING", "PROCESSING", "DISPENSING" and "CARD_EJECTED" as value.
- **TransactionType**: Enum having "WITHDRAW", "DEPOSIT", "BALANCE_INQUIRY" as value.
- **CashDenomination**: Enum having "TWO_THOUSAND", "FIVE_HUNDRED", "TWO_HUNDRED", "ONE_HUNDRED", "FIFTY", "TWENTY" and "TEN" as value.
- **Card**: Data class that holds card information like CardNumber.
- **Account**: Data class that holds Number, **AccountType**, **Card**.
- **Transaction**: Data class that holds data related to transaction like **TransactionType**, amount and timestamp.
- **CashDispenser**: 
