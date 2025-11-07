package dto

type YfcVariant struct {
	Sku             string  `json:"sku"`
	Name            string  `json:"name"`
	Price           float32 `json:"price"`
	AvailableStatus string  `json:"availableStatus"`
}

type YfcVariantCategory struct {
	Sku               string       `json:"sku"`
	Name              string       `json:"name"`
	SelectionRangeMin int32        `json:"selectionRangeMin"`
	SelectionRangeMax int32        `json:"selectionRangeMax"`
	Modifiers         []YfcVariant `json:"modifiers"`
}

type YfcMenuItem struct {
	Sku             string               `json:"sku"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	AvailableStatus string               `json:"availableStatus"`
	Price           float32              `json:"price"`
	Image           string               `json:"image"`
	Category        string               `json:"category"`
	ModifierGroups  []YfcVariantCategory `json:"modifierGroups"`
}

type YfcMenu struct {
	MerchantID string        `json:"merchantId"`
	MenuItems  []YfcMenuItem `json:"products"`
}
