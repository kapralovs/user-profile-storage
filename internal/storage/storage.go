package storage

import (
	"errors"
	"fmt"
	"log"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

func New() *Storage {
	store := &Storage{}
	return store
}

func (st *Storage) Init() {
	inMemoryStorage := make(map[string]*users.Profile, 3)
	st.Db = inMemoryStorage
	st.Db["1"] = &users.Profile{
		ID:       "1",
		Email:    "someUser@domain.com",
		Username: "SomeUer",
		Password: "simplestPassword",
		IsAdmin:  true,
	}
	st.Db["2"] = &users.Profile{
		ID:       "2",
		Email:    "johndoe@domain.com",
		Username: "john_doe",
		Password: "top123secret",
		IsAdmin:  false,
	}
	st.Db["3"] = &users.Profile{
		ID:       "3",
		Email:    "mr_robot@domain.com",
		Username: "mrR0b0T",
		Password: "anonymous",
		IsAdmin:  false,
	}
}

func (st *Storage) Load(id string) (*users.Profile, error) {
	profile, ok := st.Db[id]
	if !ok {
		log.Println("Profile loading error. A profile with this ID does not exist.")
		return nil, errors.New("it is not possible to upload a user profile because it does not exist")
	}

	log.Printf("Profile \"%s\" is loaded.\n", profile.Username)
	return profile, nil
}

func (st *Storage) Save(p *users.Profile) error {
	if p != nil {
		if p.ID == "" {
			return errors.New("can't save profile with empty ID field")
		}

		st.mu.Lock()
		defer st.mu.Unlock()
		st.Db[p.ID] = p

		log.Printf("The profile \"%s\" is saved.\n", st.Db[p.ID].Username)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}

func (st *Storage) Edit(id string, np *users.Profile) error {
	if np != nil {
		user, err := st.Load(id)
		if err != nil {
			return errors.New("it is not possible to edit a user profile because it does not exist")
		}

		fmt.Println("user loaded")
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

		if err := st.Save(user); err != nil {
			return err
		}
		return nil
	}

	return errors.New("edited user profile is nil")
}

func (st *Storage) Delete(id string) error {
	st.mu.Lock()
	defer st.mu.Unlock()
	user, err := st.Load(id)
	if err != nil {
		return errors.New("it is not possible to delete a user profile because it does not exist")
	}

	log.Printf("User profile \"%s\" has been deleted.\n", user.Username)
	delete(st.Db, user.ID)
	return nil
}

func (st *Storage) CheckForDuplicates(p *users.Profile) error {
	if p != nil {
		if p.ID != "" {
			for id, profile := range st.Db {
				if id == p.ID {
					return errors.New("profile with this ID already exists")
				}
				if profile.Username == p.Username {
					return errors.New("profile with this username already exists")
				}
				if profile.Email == p.Email {
					return errors.New("profile with this email already exists")
				}

				return nil
			}

			log.Println()
			return errors.New("empty profile ID")
		}
	}

	return errors.New("edited user profile is nil")
}
