package dto

import "time"

type Promotion struct {
	MerchantCode *string    `json:"merchant_code"`
	Name         *string    `json:"name"`
	Sku          *string    `json:"sku"`
	Price        *float32   `json:"price"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
}
