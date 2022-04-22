package server

import "net/http"

func (s *Server) Run() error {

	initRouter()
	return http.ListenAndServe(":8080", nil)
}

func initRouter() {

}
