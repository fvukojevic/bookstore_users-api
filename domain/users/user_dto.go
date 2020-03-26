/**
User Data Transfer Object
Here we create User transfer object that we can move from Persistence to Business layer
*/
package users

import (
	"github.com/fvukojevic/bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.LastName = strings.TrimSpace(user.LastName)
	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid user password")
	}

	return nil
}
