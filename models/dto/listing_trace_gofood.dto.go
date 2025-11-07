package dto

type GoFoodTraceMenu struct {
	ID             string `json:"id"`
	ExternalMenuID string `json:"external_menu_id"`
}

type GoFoodTraceOutlet struct {
	ID               string `json:"id"`
	ExternalOutletID string `json:"external_outlet_id"`
}

type GoFoodTraceBody struct {
	RequestID string            `json:"request_id"`
	Outlet    GoFoodTraceOutlet `json:"outlet"`
	Menus     []GoFoodTraceMenu `json:"menus"`
}

type GoFoodTraceHeader struct {
	EventName string `json:"event_name"`
	EventID   string `json:"event_id"`
	Version   int    `json:"version"`
	Timestamp string `json:"timestamp"`
}

type GoFoodTrace struct {
	Header GoFoodTraceHeader `json:"header"`
	Body   GoFoodTraceBody   `json:"body"`
}
