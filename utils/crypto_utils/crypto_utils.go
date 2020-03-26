package crypto_utils

import (
	"encoding/hex"
	"github.com/fvukojevic/bookstore_users-api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func GetBcrypt(input string) (*string, *errors.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), 3)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	stringHash := hex.EncodeToString(hash)
	return &stringHash, nil
}
