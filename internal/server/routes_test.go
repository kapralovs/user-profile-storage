package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func Test_create(t *testing.T) {
	// Создаем БД
	strg := storage.New()
	// заполняем БД
	strg.Init()
	/*
		ДАЛЕЕ СМОТРЕТЬ ВИДЕО GOPHER SCHOOL ПРО ТЕСТИРОВАНИЕ ВЕБ СЕРВИСОВ!!!
	*/

	type args struct {
		st *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "OkCase",
			args: args{
				st: storage.New(),
			},
			want: []byte("User profile is created!"),
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// конвертим в JSON структуру
			jsonAsBytes, _ := json.Marshal(&users.Profile{ID: "4", Email: "testNickname@testdomain.com", Username: "testNicName", Password: "TestPassword"})
			// создаем запрос на создание пользователя
			req := httptest.NewRequest("POST", "/user", strings.NewReader(string(jsonAsBytes)))
			// создаем рекордер ответа
			resp := httptest.NewRecorder()

			// создаем handler
			handler := http.HandlerFunc(create(strg))
			// вызываем create(st *storage.Storage)
			handler.ServeHTTP(resp, req)
			if reflect.DeepEqual(tt.want, resp.Body.Bytes()) {
				t.Errorf("Response from create(st) = %v, want %v", resp.Body.Bytes(), tt.want)
			}
			if got := create(tt.args.st); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("create() = %v, want %v", got, tt.want)
			}
		})
	}
}
