package response

import "swiftDaily_myself/model/database"

type Login struct {
	User              database.User `json:"user"`
	AccessToken       string        `json:"access_token"`
	AccessTokenExpire int64         `json:"access_token_expire"`
}
