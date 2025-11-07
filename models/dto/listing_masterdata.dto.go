package dto

type MasterDataModifier struct {
	ID              int     `json:"id"`
	SKU             string  `json:"sku"`
	MarketingName   string  `json:"marketingName"`
	AvailableStatus string  `json:"availableStatus"`
	Price           float32 `json:"price"`
}

type MasterDataModifierGroup struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	MarketingName     string  `json:"marketingName"`
	AvailableStatus   string  `json:"availableStatus"`
	SelectionRangeMin float32 `json:"selectionRangeMin"`
	SelectionRangeMax float32 `json:"SelectionRangeMax"`
	Modifiers         []MasterDataModifier
}

type MasterDataItem struct {
	ID              int                       `json:"id"`
	SKU             string                    `json:"sku"`
	Name            string                    `json:"name"`
	Sequence        float32                   `json:"sequence"`
	AvailableStatus string                    `json:"availableStatus"`
	Price           float32                   `json:"price"`
	Description     string                    `json:"description"`
	ModifierGroups  []MasterDataModifierGroup `json:"modifierGroups"`
}

type MasterDataCategory struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	Sequence        float32               `json:"sequence"`
	AvailableStatus string                `json:"availableStatus"`
	Items           []MasterDataItem      `json:"items"`
	ServiceHours    MasterDataServiceHour `json:"serviceHours"`
}

type MasterDataServicePeriod struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type MasterDataServiceDay struct {
	OpenPeriodType string                    `json:"openPeriodType"`
	Periods        []MasterDataServicePeriod `json:"periods"`
}

type MasterDataServiceHour struct {
	Sun MasterDataServiceDay `json:"sun"`
	Mon MasterDataServiceDay `json:"mon"`
	Tue MasterDataServiceDay `json:"tue"`
	Wed MasterDataServiceDay `json:"wed"`
	Thu MasterDataServiceDay `json:"thu"`
	Fri MasterDataServiceDay `json:"fri"`
	Sat MasterDataServiceDay `json:"sat"`
}

type MasterDataSection struct {
	ID         string               `json:"id"`
	Name       string               `json:"name"`
	Categories []MasterDataCategory `json:"categories"`
}

type MasterDataListing struct {
	MerchantID        string              `json:"merchantID"`
	PartnerMerchantID string              `json:"partnerMerchantID"`
	BrandOutletName   string              `json:"brandOutletName"`
	BrandID           int                 `json:"brandId"`
	BrandName         string              `json:"brandName"`
	WarehouseID       int                 `json:"warehouseId"`
	WarehouseName     string              `json:"warehouseName"`
	Sections          []MasterDataSection `json:"sections"`
}
