package server

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func checkCredentials(st storage.Storage, encodedCreds string) (*users.Profile, error) {
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

func authorization(st storage.Storage, w http.ResponseWriter, r *http.Request) (*users.Profile, error) {
	encodedCreds := r.Header.Get("Authorization")
	fmt.Printf("Authorization header value: %v\n", encodedCreds)
	user, err := checkCredentials(st, encodedCreds)
	if err != nil {
		log.Println(err)
		w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
		w.WriteHeader(401)

		return nil, err
	}

	return user, nil
}