package database

import "swiftDaily_myself/global"

type Role struct {
	global.Model
	Name string `json:"name"`
}
