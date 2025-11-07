package dto

import "time"

type Merchant struct {
	MerchantCode        string    `json:"merchant_code"`
	PartnerMerchantCode string    `json:"partner_merchant_code"`
	Name                string    `json:"name"`
	CurrencyCode        string    `json:"currency_code"`
	CurrencySymbol      string    `json:"currency_symbol"`
	CurrencyExponent    int       `json:"currency_exponent"`
	StartOperations     time.Time `json:"start_operations"`
	EndOperations       time.Time `json:"end_operations"`
	BrandCode           string    `json:"brand_code"`
	BrandName           string    `json:"brand_name"`
	OutletCode          string    `json:"outlet_code"`
	OutletName          string    `json:"outlet_name"`
	SalesChannels       []string  `json:"sales_channels"`
	IsImport            bool      `json:"is_import"`
}
