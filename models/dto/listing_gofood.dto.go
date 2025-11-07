package dto

type GoFoodVariant struct {
	ExternalID string  `json:"external_id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	InStock    bool    `json:"in_stock"`
}

type GoFoodSelection struct {
	MinQuantity int32 `json:"min_quantity"`
	MaxQuantity int32 `json:"max_quantity"`
}

type GoFoodRule struct {
	Selection GoFoodSelection `json:"selection"`
}

type GoFoodVariantCategory struct {
	ExternalID   string          `json:"external_id"`
	InternalName string          `json:"internal_name"`
	Name         string          `json:"name"`
	Rules        GoFoodRule      `json:"rules"`
	Variants     []GoFoodVariant `json:"variants"`
}

type GoFoodPeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type GoFoodOperationalHour struct {
	Sunday    []GoFoodPeriod `json:"sunday"`
	Monday    []GoFoodPeriod `json:"monday"`
	Tuesday   []GoFoodPeriod `json:"tuesday"`
	Wednesday []GoFoodPeriod `json:"wednesday"`
	Thursday  []GoFoodPeriod `json:"thursday"`
	Friday    []GoFoodPeriod `json:"friday"`
	Saturday  []GoFoodPeriod `json:"saturday"`
}

type GoFoodMenuItem struct {
	ExternalID                 string                `json:"external_id"`
	Name                       string                `json:"name"`
	Description                string                `json:"description"`
	InStock                    bool                  `json:"in_stock"`
	Price                      float32               `json:"price"`
	Image                      string                `json:"image"`
	OperationalHours           GoFoodOperationalHour `json:"operational_hours"`
	VariantCategoryExternalIDs []string              `json:"variant_category_external_ids"`
}

type GoFoodMenu struct {
	Name      string           `json:"name"`
	MenuItems []GoFoodMenuItem `json:"menu_items"`
}

type GoFoodListing struct {
	RequestID         string                  `json:"request_id"`
	Menus             []GoFoodMenu            `json:"menus"`
	VariantCategories []GoFoodVariantCategory `json:"variant_categories"`
}

type GoFoodListingHeader struct {
	MerchantID    string        `json:"merchant_id"`
	SalesChannels []string      `json:"sales_channels"`
	Listing       GoFoodListing `json:"listing"`
}
