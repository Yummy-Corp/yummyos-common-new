package dto

type Promotions struct {
	MerchantCode string  `json:"merchant_code"`
	Name         string  `json:"name"`
	Sku          string  `json:"sku"`
	Price        float32 `json:"price"`
	StartDate    int64   `json:"start_date"`
	EndDate      int64   `json:"end_date"`
}

type YfcPromotion struct {
	MerchantID string       `json:"merchantId"`
	Promotions []Promotions `json:"promotions"`
}
