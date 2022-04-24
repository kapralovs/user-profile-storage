package users

import (
	"fmt"
	"log"
)

func CheckAdminRights(profile *Profile) error {
	if !profile.IsAdmin {
		log.Printf("The administrator rights check failed. User \"%s\" is not an administrator\n", profile.Username)
		return fmt.Errorf("user \"%s\" does not have administrator rights", profile.Username)
	}

	log.Printf("The administrator rights check has been passed. User \"%s\" is administrator\n", profile.Username)
	return nil
}
