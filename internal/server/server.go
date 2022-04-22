package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func New() *Server {
	return &Server{
		router:  mux.NewRouter(),
		Storage: storage.New(),
	}
}

func (s *Server) Run() error {

	initRouter()
	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) initRouter() {
	s.router.HandleFunc("/user", getListOfUsers).Methods("GET")
	s.router.HandleFunc("/user/{id}", getUserByID(s)).Methods("GET")
	s.router.HandleFunc("/user", createUser(s)).Methods("POST")
}

func editProfile(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func getListOfUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string())
}

func getUserByID(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(s.Storage)
	}
}

func createUser(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser users.UserProfile
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Fatal(err)
		}

		s.Storage[newUser.ID] = &newUser
	}
}
