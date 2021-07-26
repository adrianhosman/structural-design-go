package usecase

import (
	"github.com/adrianhosman/structural-design-go/model"
)

//Calculate invoice for provide the data on invoice
func (u *impl) CalculateInvoiceData(businessID string) (*model.CalculationInvoiceResponse, error) {
	var (
		result = new(model.CalculationInvoiceResponse)
	)

	return result, nil
}
