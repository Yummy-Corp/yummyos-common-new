package dto

type GoFoodListingRespondError struct {
	MessageTitle string                 `json:"message_title"`
	Message      string                 `json:"message"`
	Data         map[string]interface{} `json:"data"`
}

type GoFoodListingRespond struct {
	Success bool                        `json:"success"`
	Errors  []GoFoodListingRespondError `json:"errors"`
}
