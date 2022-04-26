package server

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
)

type Server struct {
	router  *mux.Router
	storage *storage.Storage
}
