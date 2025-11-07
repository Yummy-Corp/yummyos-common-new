package dto

type PriceMerchant struct {
	MerchantCode string  `json:"merchant_code"`
	Price        float32 `json:"price"`
}

type Price struct {
	Type           string          `json:"type"`
	Code           string          `json:"code"`
	Name           string          `json:"name"`
	PriceMerchants []PriceMerchant `json:"price_merchants"`
}
