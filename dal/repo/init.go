package repo

import (
	"github.com/adrianhosman/structural-design-go/domain"
)

type InvoiceDAL interface {
	GetDeposits(businessID string, fileName string) (map[string][]*domain.Deposits, error)
	GetVAData(businessID string, fileName string) ([]*domain.VAData, error)
	GetBillingRatesData(businessID string, fileName string) (*domain.BillingRates, error)
	GetBankAccountRequested(businessID string, fileName string) ([]*domain.BankAccountRequest, error)
}

type impl struct {
}

func New() InvoiceDAL {
	return &impl{}
}
