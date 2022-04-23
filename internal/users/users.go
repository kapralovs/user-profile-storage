package users

import (
	"fmt"
)

// func Create() error {
// 	return nil
// }

// func Delete() error {
// 	return nil
// }

// func Edit() error {
// 	return nil
// }

// func Load() (*UserProfile, error) {
// 	return usr, nil
// }

func CheckAdminRights(profile *UserProfile) error {
	if !profile.IsAdmin {
		return fmt.Errorf("user \"%s\" does not have administrator rights", profile.Username)
	}

	return nil
}
