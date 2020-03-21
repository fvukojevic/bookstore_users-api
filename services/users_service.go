/**
Entire business logic goes here.
*/
package services

import (
	"github.com/fvukojevic/bookstore_users-api/domain/users"
	"github.com/fvukojevic/bookstore_users-api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}

	return user, nil
}
