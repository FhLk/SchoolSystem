package models

import "time"

type Member struct {
	Id        string    `gorm:"column:Id;primary_key" json:"id"`
	Firstname string    `gorm:"column:First_name" json:"firstname"`
	Lastname  string    `gorm:"column:Last_name" json:"lastname"`
	Class     string    `gorm:"column:Class" json:"class"`
	EXP       time.Time `gorm:"column:Exp_card" json:"exp"`
}
