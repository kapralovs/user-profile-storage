package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
)

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) Run() error {
	st := storage.New()
	st.Init()
	s.storage = st
	s.initRouter()
	log.Println("Server is starting on port :8080")
	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) initRouter() {
	s.router.HandleFunc("/user", getProfiles(s.storage)).Methods("GET")
	s.router.HandleFunc("/user", create(s.storage)).Methods("POST")
	s.router.HandleFunc("/user/{id}", getProfileByID(s.storage)).Methods("GET")
	s.router.HandleFunc("/user/{id}", edit(s.storage)).Methods("POST")
	s.router.HandleFunc("/user/{id}", remove(s.storage)).Methods("DELETE")
}
