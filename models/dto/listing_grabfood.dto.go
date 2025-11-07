package dto

type GrabFoodModifier struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	AvailableStatus string  `json:"availableStatus"`
	Price           float32 `json:"price"`
}

type GrabFoodModifierGroup struct {
	ID                string             `json:"id"`
	Name              string             `json:"name"`
	AvailableStatus   string             `json:"availableStatus"`
	SelectionRangeMin int32              `json:"selectionRangeMin"`
	SelectionRangeMax int32              `json:"selectionRangeMax"`
	Modifiers         []GrabFoodModifier `json:"modifiers"`
}

type GrabFoodItem struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"name"`
	SellingTimeID   string                  `json:"sellingTimeID"`
	AvailableStatus string                  `json:"availableStatus"`
	Description     string                  `json:"description"`
	Price           float32                 `json:"price"`
	Photos          []string                `json:"photos"`
	ModifierGroups  []GrabFoodModifierGroup `json:"modifierGroups"`
}

type GrabFoodCategory struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	SellingTimeID   string         `json:"sellingTimeID"`
	AvailableStatus string         `json:"availableStatus"`
	Items           []GrabFoodItem `json:"items"`
}

type GrabFoodServicePeriod struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type GrabFoodServiceDay struct {
	OpenPeriodType string                  `json:"openPeriodType"`
	Periods        []GrabFoodServicePeriod `json:"periods"`
}

type GrabFoodServiceHour struct {
	Sun GrabFoodServiceDay `json:"sun"`
	Mon GrabFoodServiceDay `json:"mon"`
	Tue GrabFoodServiceDay `json:"tue"`
	Wed GrabFoodServiceDay `json:"wed"`
	Thu GrabFoodServiceDay `json:"thu"`
	Fri GrabFoodServiceDay `json:"fri"`
	Sat GrabFoodServiceDay `json:"sat"`
}

type GrabFoodSellingTime struct {
	StartTime    string              `json:"startTime"`
	EndTime      string              `json:"endTime"`
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	ServiceHours GrabFoodServiceHour `json:"serviceHours"`
}

type GrabFoodCurrency struct {
	Code     string `json:"code"`
	Exponent int32  `json:"exponent"`
	Symbol   string `json:"symbol"`
}

type GrabFoodListing struct {
	MerchantID        string                `json:"merchantID"`
	PartnerMerchantID string                `json:"partnerMerchantID"`
	Currency          GrabFoodCurrency      `json:"currency"`
	SellingTimes      []GrabFoodSellingTime `json:"sellingTimes"`
	Categories        []GrabFoodCategory    `json:"categories"`
}

type GrabFoodListingHeader struct {
	MerchantID    string          `json:"merchant_id"`
	SalesChannels []string        `json:"sales_channels"`
	Listing       GrabFoodListing `json:"listing"`
}
