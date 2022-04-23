package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
)

func New(st storage.Storage) *Server {
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
