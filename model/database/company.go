package database

import "swiftDaily_myself/global"

type Company struct {
	global.Model
	Address     string `json:"address"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	UserID      uint   `json:"user_id"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
