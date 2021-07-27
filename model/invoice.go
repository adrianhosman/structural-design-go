package model

type CalculationInvoiceResponse struct {
	BusinessID             string               `json:"business_id"`
	SubtotalPrice          float64              `json:"subtotal"`
	SubTotalPriceFormatted string               `json:"subtotal_price_formatted"`
	DetailProduct          []*DetailProductData `json:"detail_product"`
	OtherProduct           []*DetailProductData `json:"other_product"`
}

type DetailProductData struct {
	Name                  string `json:"name"`
	UnitPriceFormatted    string `json:"unit_price_formatted"`
	Qty                   int    `json:"qty"`
	VolumeFormatted       string `json:"volume_formatted"`
	BilledAmountFormatted string `json:"billed_amount_formatted"`
}
