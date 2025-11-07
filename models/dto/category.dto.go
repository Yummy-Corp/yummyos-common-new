package dto

type Category struct {
	CategoryCode string `json:"category_code"`
	Rank         int    `json:"rank"`
	Name         string `json:"name"`
	ParentCode   string `json:"parent_code"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}
