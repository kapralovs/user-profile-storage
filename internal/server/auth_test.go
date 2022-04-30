package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func Test_checkCredentials(t *testing.T) {
	type args struct {
		st           *storage.Storage
		encodedCreds string
	}
	tests := []struct {
		name    string
		args    args
		want    *users.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkCredentials(tt.args.st, tt.args.encodedCreds)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorization(t *testing.T) {
	type args struct {
		st *storage.Storage
		w  http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *users.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authorization(tt.args.st, tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorization() = %v, want %v", got, tt.want)
			}
		})
	}
}
