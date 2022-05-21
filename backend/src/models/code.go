package models

type Code struct{
	ID uint `json:"id" gorm:"primary_key"`
	Code string `json:"code"`
}