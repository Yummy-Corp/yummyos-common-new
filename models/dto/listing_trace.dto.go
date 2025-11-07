package dto

type ListingTrace struct {
	RequestID  string   `json:"requestID"`
	MerchantID string   `json:"merchantID"`
	JobID      string   `json:"jobID"`
	UpdatedAt  string   `json:"updatedAt"`
	Status     string   `json:"status"`
	Errors     []string `json:"errors"`
}
