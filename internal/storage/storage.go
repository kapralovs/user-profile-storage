package storage

import "github.com/kapralovs/user-profile-storage/internal/users"

func New() Storage {
	store := make(Storage)
	return store
}

func (s Storage) InitStorage() {
	s["1"] = &users.UserProfile{
		ID:       "1",
		Email:    "user1@domain.com",
		Username: "user1",
		Password: "password1",
		IsAdmin:  false,
	}
	s["2"] = &users.UserProfile{
		ID:       "2",
		Email:    "user2@domain.com",
		Username: "user2",
		Password: "password2",
		IsAdmin:  false,
	}
	s["3"] = &users.UserProfile{
		ID:       "3",
		Email:    "user3@domain.com",
		Username: "user3",
		Password: "password3",
		IsAdmin:  false,
	}
}
