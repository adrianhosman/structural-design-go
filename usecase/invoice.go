package usecase

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/adrianhosman/structural-design-go/model"
	"github.com/adrianhosman/structural-design-go/util"
)

//Calculate invoice for provide the data on invoice
func (u *impl) CalculateInvoiceData(businessID string) (*model.CalculationInvoiceResponse, error) {
	var (
		result             = new(model.CalculationInvoiceResponse)
		vaAmount           float64
		mapEwallet         = make(map[string]float64, 0)
		bankAccountRequest float64
	)
	result.DetailProduct = make([]*model.DetailProductData, 0)

	fileNameVa := fmt.Sprintf("VAData_%s.csv", businessID)
	vaData, err := u.repo.GetVAData(businessID, fileNameVa)
	if err != nil {
		return nil, err
	}

	for _, dt := range vaData {
		vaAmount += dt.Amount
	}

	fileNameBillingRates := fmt.Sprintf("BillingRates_%s.csv", businessID)
	billingRates, err := u.repo.GetBillingRatesData(businessID, fileNameBillingRates)
	if err != nil {
		return nil, err
	}

	if billingRates != nil {
		vaTemp := strings.Replace(billingRates.VirtualAccountRates, " ", "", -1)
		vaTemp = strings.Replace(vaTemp, "[", "", -1)
		vaTemp = strings.Replace(vaTemp, "]", "", -1)

		t := regexp.MustCompile(`[,\]\[]`)
		temp := t.Split(vaTemp, -1)
		min := float64(1000000)
		bankAccountRequest = billingRates.BankAccountRequested

		ewTemp := strings.Replace(billingRates.EwalletRates, " ", "", -1)
		ewTemp = strings.Replace(ewTemp, "[", "", -1)
		ewTemp = strings.Replace(ewTemp, "]", "", -1)
		eWalletTemp := t.Split(ewTemp, -1)
		ewRes := make([]string, 0)
		for _, res := range eWalletTemp {
			ewTemp := strings.Replace(res, " ", "", -1)
			ewTemp = strings.Replace(ewTemp, "[", "", -1)
			ewTemp = strings.Replace(ewTemp, "]", "", -1)
			ewRes = append(ewRes, ewTemp)
		}
		for i, _ := range ewRes {
			if (i+1)%3 == 0 && i != 0 {
				tes := strings.Replace(ewRes[i], "u", "", -1)
				tes = strings.Replace(tes, "'", "", -1)

				mapEwallet[tes], _ = strconv.ParseFloat(ewRes[i-1], 64)
			}
		}

		for i, _ := range temp {
			if i == 0 {
				continue
			}
			price, _ := strconv.ParseFloat(strings.TrimSpace(temp[i]), 64)
			if min > price {
				min = price
			}
		}
		bill := min * float64(len(vaData))
		result.SubtotalPrice += bill
		result.DetailProduct = append(result.DetailProduct, &model.DetailProductData{
			Name:                  "Virtual Account",
			Qty:                   len(vaData),
			VolumeFormatted:       util.IDR(vaAmount),
			UnitPriceFormatted:    util.IDR(min),
			BilledAmountFormatted: util.IDR(bill),
		})

	}

	fileName := fmt.Sprintf("Deposits_%s.csv", businessID)
	deposits, err := u.repo.GetDeposits(businessID, fileName)
	if err != nil {
		return nil, err
	}
	if deposits == nil {
		return nil, err
	}
	var eWalletTotal float64
	for name, dt := range deposits {
		var jum float64
		for _, x := range dt {
			jum += x.TransactionAmount
		}
		resTemp := &model.DetailProductData{
			Name:            name,
			Qty:             len(dt),
			VolumeFormatted: util.IDR(jum),
		}

		if mapEwallet[name] > 0 {
			resTemp.UnitPriceFormatted = fmt.Sprintf("%.1f%%", (mapEwallet[name] * 100))
			bill := float64(len(dt)) * jum
			resTemp.BilledAmountFormatted = util.IDR(bill)
			result.SubtotalPrice = bill
		} else {
			resTemp.BilledAmountFormatted = util.IDR(jum)
			result.SubtotalPrice = jum
		}
		result.DetailProduct = append(result.DetailProduct, resTemp)

		eWalletTotal += jum
	}

	fileNameBankAccount := fmt.Sprintf("BankAccount_%s.csv", businessID)
	baRequest, err := u.repo.GetBankAccountRequested(businessID, fileNameBankAccount)
	if err != nil {
		return nil, err
	}
	if baRequest == nil {
		return nil, err
	}

	var baRequestedSum float64

	baRequestedSum = bankAccountRequest * float64(len(baRequest))
	result.SubtotalPrice += baRequestedSum
	result.OtherProduct = append(result.OtherProduct, &model.DetailProductData{
		Name:                  "Name Validator",
		Qty:                   len(baRequest),
		UnitPriceFormatted:    util.IDR(bankAccountRequest),
		BilledAmountFormatted: util.IDR(baRequestedSum),
	})

	result.BusinessID = businessID
	result.SubtotalPrice = vaAmount + eWalletTotal
	result.SubTotalPriceFormatted = util.IDR(result.SubtotalPrice)

	return result, nil
}
