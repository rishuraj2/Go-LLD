### Create Account
    POST /account
    request:
    {
        "account-type":
    }

    response:
    {
        "status":
        "message":
        "account-number":
    }

### Create Card
     POST /card
    request:
    {
        "account_number":
        "pin":
    }

    response:
    {
        "status":
        "message":
        "card_number":
    }

### Authenticate Card

    POST /card/{card_number}
    request:
    {
        "pin":
    }

    response:
    {
        "status":
        "message":
        "account_number":
    }

### Credit Amount
    POST /account/{account_number}/deposit
    request:
    {
        "amount":
    }

    response:
    {
        "status":
        "message":
        "transaction_details": {
            "id":
            "type":
            "account_number":
            "amount":
            "time_stamp":
        }
    }

### Withdraw Amount
    POST /account/{account_number}/withdraw
    request:
    {
        "amount":
    }

    response:
    {
        "status":
        "message":
        "transaction_details": {
            "id":
            "type":
            "account_number":
            "amount":
            "time_stamp":
        }
    }

### Fetch Transaction History

    GET /account/{account_number}/all_transactions
    request:
    {
        EMPTY
    }

    response:
    {
        "status":
        "message":
        "transactions": [
            {
                "id":
                "type":
                "account_number":
                "amount":
                "time_stamp":
            },
            {
                "id":
                "type":
                "account_number":
                "amount":
                "time_stamp":
            },
            ...
        ]
    }
