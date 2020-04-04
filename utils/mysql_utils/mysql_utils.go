package mysql_utils

import (
	"fmt"
	"github.com/fvukojevic/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("No record matching given Id")
		}
		return errors.NewInternalServerError("Error passing db response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid data"))
	}

	return errors.NewInternalServerError("Error processing request")
}
