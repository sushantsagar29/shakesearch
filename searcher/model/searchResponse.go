package model

type Response struct {
	Count   int      `json:"count"`
	Matches []string `json:"matches"`
}
