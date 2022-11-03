package models

type Dashboard struct {
	User   int `json:"user"`
	Total  int `json:"total"`
	Active int `json:"active"`
}
