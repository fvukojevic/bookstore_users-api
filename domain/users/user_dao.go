/**
User Access Object
The Entire logic to Persist and Retrieve user from the database
Only point in the application where we interact with database (Mysql, Cassandra, Mongo, Postgresql)
If tomorrow we change the db, we only should have to change this file. Entire app works
*/
package users

import (
	"fmt"
	"github.com/fvukojevic/bookstore_users-api/datasources/mysql/users_db"
	"github.com/fvukojevic/bookstore_users-api/utils/date_utils"
	"github.com/fvukojevic/bookstore_users-api/utils/errors"
	"strings"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	indexUnique     = "Duplicate entry"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUnique) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when tryting to save the user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when tryting to save the user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
