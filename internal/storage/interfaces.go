package storage

import (
	"github.com/kapralovs/user-profile-storage/internal/users"
)

type Storage map[string]*users.User
