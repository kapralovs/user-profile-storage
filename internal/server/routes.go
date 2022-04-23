package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/auth"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func create(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем креды из хедера
		encodedCreds := r.Header.Get("Authorization")
		fmt.Printf("Authorization header value: %v\n", encodedCreds) //Debug
		// Проверяем креды на корректность
		user, err := auth.CheckCredentials(s.storage, encodedCreds)
		if err != nil {
			// Если креды не корректные, тогда логгируем ошибку...
			log.Println(err)
			// ...затем выставляем хедер для респонса,...
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)

			// ...пишем в хедер статус код 401 - Not Authorized...
			w.WriteHeader(401)

			// ...и пишем текст ошибки в ответ
			fmt.Fprintln(w, err)

			return
		}

		// Если с кредами все ок, тогда создаем профиль под нового пользователя...
		var newUser *users.Profile

		// ...затем читаем из ответа JSON с данными нового профиля
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err) //logging error
		}

		// Парсим JSON, заполняя соответствующие поля для профиля нового пользователя
		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Fatal(err) //logging error
		}

		// И пишем профиль нового пользователя в БД
		s.storage[newUser.ID] = newUser
	}
}

func edit(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedCreds := r.Header.Get("Authorization")
		user, err := auth.CheckCredentials(s.storage, encodedCreds)
		if err != nil {
			log.Println(err)
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds) //add response header
			w.WriteHeader(401)                                              //response status code
			fmt.Fprintln(w, err)                                            //send response

			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err.Error())
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		var editedProfile *users.Profile
		if err := json.Unmarshal(body, editedProfile); err != nil {
			log.Println(err)
		}

		vars := mux.Vars(r)
		id := vars["id"]
		s.storage[id] = editedProfile
	}
}

func remove(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedCreds := r.Header.Get("Authorization")
		user, err := auth.CheckCredentials(s.storage, encodedCreds)
		if err != nil {
			log.Println(err)
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
			w.WriteHeader(401)
			fmt.Fprintln(w, err)
		}

		vars := mux.Vars(r)
		id := vars["id"]
		delete(s.storage, id)
	}
}

func getProfiles(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// s.storage["1"] = &users.UserProfile{
		// 	ID:       "1",
		// 	Email:    "user1@domain.com",
		// 	Username: "user1",
		// 	Password: "password1",
		// 	IsAdmin:  false,
		// }
		// s.storage["2"] = &users.UserProfile{
		// 	ID:       "2",
		// 	Email:    "user2@domain.com",
		// 	Username: "user2",
		// 	Password: "password2",
		// 	IsAdmin:  false,
		// }
		// s.storage["3"] = &users.UserProfile{
		// 	ID:       "3",
		// 	Email:    "user3@domain.com",
		// 	Username: "user3",
		// 	Password: "password3",
		// 	IsAdmin:  false,
		// }

		var profiles []*users.Profile
		for id := range s.storage {
			profiles = append(profiles, s.storage[id])
		}

		json.NewEncoder(w).Encode(profiles)
	}
}

func getProfileByID(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewEncoder(w).Encode(s.storage[id])
	}
}
