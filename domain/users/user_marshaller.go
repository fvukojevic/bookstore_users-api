/**
How our domain is going to be presented to the user
*/

package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		_ = json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, 0)
	for _, user := range users {
		result = append(result, user.Marshall(isPublic))
	}
	return result
}
