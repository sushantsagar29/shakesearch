package model

type Request struct {
	Query           string
	IsCaseSensitive bool
	IsExactMatch    bool
}
