package storage

import (
	"errors"
	"log"

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
	profile, ok := st[id]
	if ok {
		log.Printf("Profile \"%s\" is loaded.\n", profile.Username)
		return profile, nil
	}

	log.Printf("Profile loading error. The profile \"%s\" does not exist.\n", profile.Username)
	return nil, errors.New("it is not possible to upload a user profile because it does not exist")
}

func (st Storage) Save(p *users.Profile) {
	st[p.ID] = p

	log.Printf("The profile \"%s\" is saved.\n", st[p.ID].Username)
}

func (st Storage) Edit(id string, np *users.Profile) error {
	user, err := st.Load(id)
	if err != nil {
		return errors.New("it is not possible to edit a user profile because it does not exist")
	}

	if user.Email != np.Email {
		user.Email = np.Email
		log.Printf("Profile \"%s\" is edited (email)", user.Username)
	}

	if user.Username != np.Username {
		user.Username = np.Username
		log.Printf("Profile \"%s\" is edited (user)", user.Username)
	}

	if user.Password != np.Password {
		user.Password = np.Password
		log.Printf("Profile \"%s\" is edited (password)", user.Username)
	}

	if user.IsAdmin != np.IsAdmin {
		user.IsAdmin = np.IsAdmin
		log.Printf("Profile \"%s\" now has admin rights", user.Username)
	}

	st.Save(user)
	return nil
}

func (st Storage) Delete(id string) error {
	user, err := st.Load(id)
	if err != nil {
		return errors.New("it is not possible to delete a user profile because it does not exist")
	}

	log.Printf("User profile \"%s\" has been deleted.\n", user.Username)
	delete(st, user.ID)
	return nil
}

func (st Storage) CheckForDuplicates(p *users.Profile) error {
	for id, profile := range st {
		if id == p.ID {
			return errors.New("profile with this ID already exists")
		}
		if profile.Username == p.Username {
			return errors.New("profile with this username already exists")
		}
		if profile.Email == p.Email {
			return errors.New("profile with this email already exists")
		}
	}

	return nil
}
