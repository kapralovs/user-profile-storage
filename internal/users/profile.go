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
