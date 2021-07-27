package domain

type Deposits struct {
	ID                string  `json:"id"`
	BusinessID        string  `json:"business_id"`
	TransactionAmount float64 `json:"transaction_amount"`
	Currency          string  `json:"currency"`
	EwalletType       string  `json:"ewallet_type"`
	BillingRates      int     `json:"rates"`
	BillingChargeType string  `json:"billing_charge_type"`
}

type VAData struct {
	Amount float64 `json:"amount"`
}

type BillingRates struct {
	ID                   string  ` json:"ID"`
	BusinessID           string  `json:"business_id"`
	BankAccountRequested float64 `json:"bank_account_requested"`
	EwalletRates         string  `json:"ewallet_rates"`
	VirtualAccountRates  string  `json:"virtual_account_rates"`
}

type BankAccountRequest struct {
	ID string `json:"id"`
}
