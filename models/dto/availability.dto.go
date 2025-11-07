package dto

type Availability struct {
	Type            string `json:"type"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	AvailableStatus string `json:"available_status"`
}

type Availabilities struct {
	Availabilities []Availability `json:"availabilities"`
}
