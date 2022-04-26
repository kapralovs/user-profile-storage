package storage

import (
	"sync"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

type Storage struct {
	mu sync.Mutex
	Db map[string]*users.Profile
}
