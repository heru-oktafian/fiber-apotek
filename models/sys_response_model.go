package models

// Response represents the standard JSON response format / structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response Get All represents the standard JSON response format / structure
type ResponseGetAll struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Search      string      `json:"search"`
	TotalItems  int         `json:"total_items"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	PerPage     int         `json:"per_page"`
	Data        interface{} `json:"data"`
}
