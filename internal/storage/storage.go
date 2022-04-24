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
		Email:    "someUser@domain.com",
		Username: "SomeUer",
		Password: "simplestPassword",
		IsAdmin:  true,
	}
	s["2"] = &users.Profile{
		ID:       "2",
		Email:    "johndoe@domain.com",
		Username: "john_doe",
		Password: "top123secret",
		IsAdmin:  false,
	}
	s["3"] = &users.Profile{
		ID:       "3",
		Email:    "mr_robot@domain.com",
		Username: "mrR0b0T",
		Password: "anonymous",
		IsAdmin:  false,
	}
}

func (st Storage) Load(id string) (*users.Profile, error) {
	if profile, ok := st[id]; ok {
		return profile, nil
	}

	return nil, errors.New("it is not possible to upload a user profile because it does not exist")
}

func (st Storage) Save(p *users.Profile) error {
	if _, ok := st[p.ID]; ok {
		return errors.New("user with this ID is already exists")
	}

	st[p.ID] = p

	return nil
}

func (st Storage) Edit(id string, np *users.Profile) error {
	user, err := st.Load(id)
	if err != nil {
		return errors.New("it is not possible to edit a user profile because it does not exist")
	}

	if user.Email != np.Email {
		user.Email = np.Email
	}

	if user.Username != np.Username {
		user.Username = np.Username
	}

	if user.Password != np.Password {
		user.Password = np.Password
	}

	if user.IsAdmin != np.IsAdmin {
		user.IsAdmin = np.IsAdmin
	}

	st.Save(user)
	return nil
}

func (st Storage) Delete(id string) error {
	user, err := st.Load(id)
	if err != nil {
		return errors.New("it is not possible to delete a user profile because it does not exist")
	}

	delete(st, user.ID)
	return nil
}
