package dto

type OdooModifier struct {
	ID              int     `json:"id"`
	SKU             string  `json:"sku"`
	MarketingName   string  `json:"marketingName"`
	AvailableStatus string  `json:"availableStatus"`
	Price           float32 `json:"price"`
}

type OdooModifierGroup struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	MarketingName     string  `json:"marketingName"`
	AvailableStatus   string  `json:"availableStatus"`
	SelectionRangeMin float32 `json:"selectionRangeMin"`
	SelectionRangeMax float32 `json:"SelectionRangeMax"`
	Modifiers         []OdooModifier
}

type OdooItem struct {
	ID              int                 `json:"id"`
	SKU             string              `json:"sku"`
	Name            string              `json:"name"`
	Sequence        float32             `json:"sequence"`
	AvailableStatus string              `json:"availableStatus"`
	Price           float32             `json:"price"`
	Description     string              `json:"description"`
	ModifierGroups  []OdooModifierGroup `json:"modifierGroups"`
}

type OdooCategory struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Sequence        float32    `json:"sequence"`
	AvailableStatus string     `json:"availableStatus"`
	Items           []OdooItem `json:"items"`
}

type OdooServicePeriod struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type OdooServiceDay struct {
	OpenPeriodType string              `json:"openPeriodType"`
	Periods        []OdooServicePeriod `json:"periods"`
}

type OdooServiceHour struct {
	Sun OdooServiceDay `json:"sun"`
	Mon OdooServiceDay `json:"mon"`
	Tue OdooServiceDay `json:"tue"`
	Wed OdooServiceDay `json:"wed"`
	Thu OdooServiceDay `json:"thu"`
	Fri OdooServiceDay `json:"fri"`
	Sat OdooServiceDay `json:"sat"`
}

type OdooSection struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	ServiceHours OdooServiceHour `json:"serviceHours"`
	Categories   []OdooCategory  `json:"categories"`
}

type OdooListing struct {
	MerchantID        string        `json:"merchantID"`
	PartnerMerchantID string        `json:"partnerMerchantID"`
	BrandOutletName   string        `json:"brandOutletName"`
	BrandID           int           `json:"brandId"`
	BrandName         string        `json:"brandName"`
	WarehouseID       int           `json:"warehouseId"`
	WarehouseName     string        `json:"warehouseName"`
	Sections          []OdooSection `json:"sections"`
}
