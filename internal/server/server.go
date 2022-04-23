package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/auth"
	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
		// storage: storage.New(),
	}
}

func (s *Server) Run() error {
	storage := storage.New()
	storage.InitStorage()
	s.storage = storage
	s.initRouter()
	log.Println("Server is starting on port :8080")
	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) initRouter() {
	s.router.HandleFunc("/user", getProfiles(s)).Methods("GET")
	s.router.HandleFunc("/user", create(s)).Methods("POST")
	s.router.HandleFunc("/user/{id}", getProfileByID(s)).Methods("GET")
	s.router.HandleFunc("/user/{id}", edit(s)).Methods("POST")
	s.router.HandleFunc("/user/{id}", remove(s)).Methods("DELETE")
}

func create(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedCreds := r.Header.Get("Authorization")
		if err := auth.CheckCredentials(s.storage, encodedCreds); err != nil {
			log.Println(err)
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
			w.WriteHeader(401)
			fmt.Fprintln(w, err)
		}
		var newUser *users.UserProfile
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Fatal(err)
		}

		s.storage[newUser.ID] = newUser
	}
}

func edit(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedCreds := r.Header.Get("Authorization")
		if err := auth.CheckCredentials(s.storage, encodedCreds); err != nil {
			log.Println(err)
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
			w.WriteHeader(401)
			fmt.Fprintln(w, err)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var editedProfile *users.UserProfile
		if err := json.Unmarshal(body, editedProfile); err != nil {
			log.Fatal(err)
		}

		vars := mux.Vars(r)
		id := vars["id"]
		s.storage[id] = editedProfile
	}
}

func remove(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedCreds := r.Header.Get("Authorization")
		if err := auth.CheckCredentials(s.storage, encodedCreds); err != nil {
			log.Println(err)
			w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
			w.WriteHeader(401)
			fmt.Fprintln(w, err)
		}

		// id := r.URL.Path[len("/user/"):]
		id := r.URL.Query().Get("id")
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

		var profiles []*users.UserProfile
		for id := range s.storage {
			profiles = append(profiles, s.storage[id])
		}

		json.NewEncoder(w).Encode(profiles)
		// fmt.Fprintln(w)
	}
}

func getProfileByID(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewEncoder(w).Encode(s.storage[id])
	}
}
