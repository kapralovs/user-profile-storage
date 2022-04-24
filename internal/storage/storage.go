package storage

import (
	"errors"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

func New() Storage {
	store := make(Storage)
	return store
}

func (s Storage) Init() {
	s["1"] = &users.Profile{
		ID:       "1",
		Email:    "user1@domain.com",
		Username: "user1",
		Password: "password1",
		IsAdmin:  false,
	}
	s["2"] = &users.Profile{
		ID:       "2",
		Email:    "user2@domain.com",
		Username: "user2",
		Password: "password2",
		IsAdmin:  false,
	}
	s["3"] = &users.Profile{
		ID:       "3",
		Email:    "user3@domain.com",
		Username: "user3",
		Password: "password3",
		IsAdmin:  false,
	}
}

func (st Storage) LoadProfile(id string) *users.Profile {
	if profile, ok := st[id]; ok {
		return profile
	}

	return nil
}

func (st Storage) SaveProfile(p *users.Profile) error {
	if _, ok := st[p.ID]; ok {
		return errors.New("user with this ID is already exists")
	}

	st[p.ID] = p

	return nil
}
