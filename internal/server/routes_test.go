package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func Test_create(t *testing.T) {
	var (
		strg        = storage.New()
		jsonAsBytes []byte
		req         *http.Request
	)
	strg.Init()
	validCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	notValidCreds := "RandomNotValidCreds"
	notAdminCreds := "am9obl9kb2U6dG9wMTIzc2VjcmV0"

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK_Case",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("User profile is created!"),
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("authorization failed"),
		},
		{
			name: "Without_Admin_Rights",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("user \"john_doe\" does not have administrator rights"),
		},
		{
			name: "Not_Valid_JSON",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("invalid character ',' after top-level value"),
		},
		{
			name: "Empty_Request_Body",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("unexpected end of JSON input"),
		},
		// {
		// 	// Empty storage case
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "OK_Case":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "4", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword"})
				validReq := httptest.NewRequest("POST", "/user", strings.NewReader(string(jsonAsBytes)))
				validReq.Header.Add("Authorization", "Basic "+validCreds)
				req = validReq
			case "Not_Authorized":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "4", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword"})
				notValidReq := httptest.NewRequest("POST", "/user", strings.NewReader(string(jsonAsBytes)))
				notValidReq.Header.Add("Authorization", "Basic "+notValidCreds)
				req = notValidReq
			case "Without_Admin_Rights":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "4", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword"})
				notValidReq := httptest.NewRequest("POST", "/user", strings.NewReader(string(jsonAsBytes)))
				notValidReq.Header.Add("Authorization", "Basic "+notAdminCreds)
				req = notValidReq
			case "Not_Valid_JSON":
				jsonAsBytes = []byte("\"testNickname@testdomain.com\",\"username\":\"testNickName\"}")
				notValidReq := httptest.NewRequest("POST", "/user", strings.NewReader(string(jsonAsBytes)))
				notValidReq.Header.Add("Authorization", "Basic "+validCreds)
				req = notValidReq
			case "Empty_Request_Body":
				notValidReq := httptest.NewRequest("POST", "/user", nil)
				notValidReq.Header.Add("Authorization", "Basic "+validCreds)
				req = notValidReq
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(create(strg))
			handler.ServeHTTP(resp, req)
			if !reflect.DeepEqual(tt.want, resp.Body.String()) {
				t.Errorf("Response from create(st) = %v, want %v", resp.Body.String(), tt.want)
			}
		})
	}
}

func Test_edit(t *testing.T) {
	var (
		strg        = storage.New()
		jsonAsBytes []byte
		req         *http.Request
	)
	strg.Init()
	validCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	notValidCreds := "RandomNotValidCreds"
	notAdminCreds := "am9obl9kb2U6dG9wMTIzc2VjcmV0"

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("User profile edited!"),
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("authorization failed"),
		},
		{
			name: "Without_Admin_Rights",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("user \"john_doe\" does not have administrator rights"),
		},
		// {
		// 	// Empty storage case
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "OK":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "234", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword", IsAdmin: false})
				newReq := func(method, path, body string, vars map[string]string) *http.Request {
					r := httptest.NewRequest(method, path, strings.NewReader(body))
					return mux.SetURLVars(r, vars)
				}
				req = newReq("POST", "/user/3", string(jsonAsBytes), map[string]string{"id": "3"})
				req.Header.Add("Authorization", "Basic "+validCreds)
			case "Not_Authorized":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "3", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword"})
				notValidReq := httptest.NewRequest("POST", "/user/3", strings.NewReader(string(jsonAsBytes)))
				notValidReq.Header.Add("Authorization", "Basic "+notValidCreds)
				req = notValidReq
			case "Without_Admin_Rights":
				jsonAsBytes, _ = json.Marshal(&users.Profile{ID: "3", Email: "testNickname@testdomain.com", Username: "testNickName", Password: "TestPassword"})
				notValidReq := httptest.NewRequest("POST", "/user/3", strings.NewReader(string(jsonAsBytes)))
				notValidReq.Header.Add("Authorization", "Basic "+notAdminCreds)
				req = notValidReq
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(edit(strg))
			handler.ServeHTTP(resp, req)
			if !reflect.DeepEqual(tt.want, resp.Body.String()) {
				t.Errorf("Response from edit(st) = %v, want %v", resp.Body.String(), tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	var (
		strg = storage.New()
		req  *http.Request
	)
	strg.Init()
	validCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	notValidCreds := "RandomNotValidCreds"
	notAdminCreds := "am9obl9kb2U6dG9wMTIzc2VjcmV0"

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("User profile is deleted!"),
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("authorization failed"),
		},
		{
			name: "Without_Admin_Rights",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("user \"john_doe\" does not have administrator rights"),
		},
		// {
		// 	// Empty storage case
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "OK":
				newReq := func(method, path string, vars map[string]string) *http.Request {
					r := httptest.NewRequest(method, path, nil)
					return mux.SetURLVars(r, vars)
				}
				req = newReq("DELETE", "/user/3", map[string]string{"id": "3"})
				req.Header.Add("Authorization", "Basic "+validCreds)
			case "Not_Authorized":
				notValidReq := httptest.NewRequest("DELETE", "/user/2", nil)
				notValidReq.Header.Add("Authorization", "Basic "+notValidCreds)
				req = notValidReq
			case "Without_Admin_Rights":
				notValidReq := httptest.NewRequest("DELETE", "/user/2", nil)
				notValidReq.Header.Add("Authorization", "Basic "+notAdminCreds)
				req = notValidReq
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(remove(strg))
			handler.ServeHTTP(resp, req)
			if !reflect.DeepEqual(resp.Body.String(), tt.want) {
				t.Errorf("Response from remove() = %v, want %v", resp.Body.String(), tt.want)
			}
		})
	}
}

func Test_getProfiles(t *testing.T) {
	var (
		strg = storage.New()
		req  *http.Request
	)
	strg.Init()
	validCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	notValidCreds := "RandomNotValidCreds"

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("{\"id\":\"1\",\"email\":\"someUser@domain.com\",\"username\":\"SomeUer\",\"password\":\"simplestPassword\",\"is_admin\":true}\n{\"id\":\"2\",\"email\":\"johndoe@domain.com\",\"username\":\"john_doe\",\"password\":\"top123secret\",\"is_admin\":false}\n{\"id\":\"3\",\"email\":\"mr_robot@domain.com\",\"username\":\"mrR0b0T\",\"password\":\"anonymous\",\"is_admin\":false}"),
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("authorization failed"),
		},
		// {
		// 	// Empty storage case
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "OK":
				validReq := httptest.NewRequest("GET", "/user", nil)
				validReq.Header.Add("Authorization", "Basic "+validCreds)
				req = validReq
			case "Not_Authorized":
				notValidReq := httptest.NewRequest("GET", "/user", nil)
				notValidReq.Header.Add("Authorization", "Basic "+notValidCreds)
				req = notValidReq
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(getProfiles(strg))
			handler.ServeHTTP(resp, req)
			if !reflect.DeepEqual(resp.Body.String(), tt.want) {
				t.Errorf("Response from getProfiles(st) = %v, want %v", resp.Body.String(), tt.want)
			}
		})
	}
}

func Test_getProfileByID(t *testing.T) {
	var (
		strg = storage.New()
		req  *http.Request
		// router      *mux.Router
	)
	strg.Init()
	validCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	notValidCreds := "RandomNotValidCreds"

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("{\"id\":\"3\",\"email\":\"mr_robot@domain.com\",\"username\":\"mrR0b0T\",\"password\":\"anonymous\",\"is_admin\":false}"),
		},
		{
			name: "Not_Authorized",
			args: args{
				st: strg,
			},
			want: fmt.Sprintln("authorization failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "OK":
				newReq := func(method, path string, vars map[string]string) *http.Request {
					r := httptest.NewRequest(method, path, nil)
					return mux.SetURLVars(r, vars)
				}
				req = newReq("GET", "/user/3", map[string]string{"id": "3"})
				req.Header.Add("Authorization", "Basic "+validCreds)
			case "Not_Authorized":
				notValidReq := httptest.NewRequest("GET", "/user/2", nil)
				notValidReq.Header.Add("Authorization", "Basic "+notValidCreds)
				req = notValidReq
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(getProfileByID(strg))
			handler.ServeHTTP(resp, req)
			if !reflect.DeepEqual(resp.Body.String(), tt.want) {
				t.Errorf("Response from getProfileByID(st) = %v, want %v", resp.Body.String(), tt.want)
			}
		})
	}
}
