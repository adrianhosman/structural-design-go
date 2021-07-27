package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/adrianhosman/structural-design-go/domain"
)

func (c *impl) GetDeposits(businessID string, fileName string) (map[string][]*domain.Deposits, error) {
	var (
		result = make(map[string][]*domain.Deposits, 0)
	)

	csvFile, err := os.Open("./source/" + fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		if i == 0 {
			continue
		}

		bilRate, _ := strconv.Atoi(line[9])
		amount, _ := strconv.ParseFloat(line[4], 64)
		result[line[14]] = append(result[line[14]], &domain.Deposits{
			ID:                line[0],
			BusinessID:        line[1],
			Currency:          line[13],
			TransactionAmount: amount,
			EwalletType:       line[14],
			BillingRates:      bilRate,
			BillingChargeType: line[10],
		})
	}
	return result, nil

}

func (c *impl) GetVAData(businessID string, fileName string) ([]*domain.VAData, error) {
	var (
		result = make([]*domain.VAData, 0)
	)
	csvFile, err := os.Open("./source/" + fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		tempFlo, _ := strconv.ParseFloat(line[0], 64)
		result = append(result, &domain.VAData{
			Amount: tempFlo,
		})
	}
	return result, nil
}

func (c *impl) GetBillingRatesData(businessID string, fileName string) (*domain.BillingRates, error) {
	var (
		result = new(domain.BillingRates)
	)
	csvFile, err := os.Open("./source/" + fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		nameValidatorRate, _ := strconv.ParseFloat(line[2], 64)
		result.ID = line[1]
		result.BusinessID = line[4]
		result.BankAccountRequested = nameValidatorRate
		result.EwalletRates = line[14]
		result.VirtualAccountRates = line[27]
	}

	return result, nil

}

func (c *impl) GetBankAccountRequested(businessID string, fileName string) ([]*domain.BankAccountRequest, error) {
	var (
		result = make([]*domain.BankAccountRequest, 0)
	)
	csvFile, err := os.Open("./source/" + fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		if i == 0 {
			continue
		}

		result = append(result, &domain.BankAccountRequest{ID: line[1]})
	}

	return result, nil

}
