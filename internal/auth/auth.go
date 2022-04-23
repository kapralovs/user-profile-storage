package auth

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func CheckCredentials(st storage.Storage, encodedCreds string) (*users.Profile, error) {
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return nil, err
	}

	creds := strings.Split(string(decodedCreds), ":")

	for _, profile := range st {
		if profile.Username == creds[0] && profile.Password == creds[1] {
			return profile, nil
		}
	}

	return nil, errors.New("authorisation failed because credentials are incorrect")
}
