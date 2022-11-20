package models

type Student struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	LastName  string  `json:"firstName"`
	FirstName string  `json:"lastName"`
	Average   float64 `json:"average"`
}
