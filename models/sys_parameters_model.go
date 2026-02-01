package models

// Struktur untuk menerima input dari body request
type RequestBody struct {
	Page   int    `json:"page"`
	Search string `json:"search"`
	Month  string `json:"month"`
}
