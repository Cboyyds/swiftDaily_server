package response

import "swiftDaily_myself/model/database"

type Login struct {
	User database.User `json:"user"`
}
