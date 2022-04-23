package main

import (
	"log"

	"github.com/kapralovs/user-profile-storage/internal/server"
)

func main() {
	s := server.New()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
