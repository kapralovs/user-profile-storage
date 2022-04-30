package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func Test_checkCredentials(t *testing.T) {
	var (
		validCreds    = "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
		notValidCreds = "RandomNotValidCreds"
		strg          = storage.New()
	)
	strg.Init()

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
		{
			name: "OK",
			args: args{
				st:           strg,
				encodedCreds: validCreds,
			},
			want: &users.Profile{
				ID:       "1",
				Email:    "someUser@domain.com",
				Username: "SomeUer",
				Password: "simplestPassword",
				IsAdmin:  true,
			},
		},
		{
			name: "Not_Valid_Creds",
			args: args{
				st:           strg,
				encodedCreds: notValidCreds,
			},
			want:    nil,
			wantErr: true,
		},
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
	var (
		validCreds    = "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
		notValidCreds = "RandomNotValidCreds"
		strg          = storage.New()
		someReq       = func(method, path, creds string) *http.Request {
			req := httptest.NewRequest(method, path, nil)
			req.Header.Add("Authorization", "Basic "+creds)
			return req
		}
	)
	strg.Init()

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
		{
			name: "OK",
			args: args{
				st: strg,
				w:  httptest.NewRecorder(),
				r:  someReq("GET", "/user", validCreds),
			},
			want: strg.Db["1"],
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
				w:  httptest.NewRecorder(),
				r:  someReq("GET", "/user", notValidCreds),
			},
			want:    nil,
			wantErr: true,
		},
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
