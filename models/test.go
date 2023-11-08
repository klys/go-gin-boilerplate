package models

type Test struct {
	ID int64 `json:"id"`
	Test string `json:"test"`
}

var Tests []Test