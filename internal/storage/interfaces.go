package storage

import (
	"sync"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

type Storage struct {
	mu sync.Mutex
	db map[string]*users.Profile
}
