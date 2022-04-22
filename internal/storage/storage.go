package storage

import "github.com/kapralovs/user-profile-storage/internal/users"

func New() map[string]*users.UserProfile {
	store := make(Storage)
	return store
}
