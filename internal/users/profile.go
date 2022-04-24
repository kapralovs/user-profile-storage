package users

import (
	"fmt"
)

func CheckAdminRights(profile *Profile) error {
	if !profile.IsAdmin {
		return fmt.Errorf("user \"%s\" does not have administrator rights", profile.Username)
	}

	return nil
}

func (p *Profile) Edit(np *Profile) {
	if p.Email != np.Email {
		p.Email = np.Email
	}

	if p.Username != np.Username {
		p.Username = np.Username
	}

	if p.Password != np.Password {
		p.Password = np.Password
	}

	if p.IsAdmin != np.IsAdmin {
		p.IsAdmin = np.IsAdmin
	}
}
