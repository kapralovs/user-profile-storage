package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/user-profile-storage/internal/storage"
	"github.com/kapralovs/user-profile-storage/internal/users"
)

func create(st storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err.Error())
			return
		}

		var newUser *users.Profile

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Fatal(err)
		}

		if err := st.SaveProfile(newUser); err != nil {
			fmt.Fprintln(w, err)
		}
	}
}

func edit(st storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err.Error())
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		var editedProfile *users.Profile
		if err := json.Unmarshal(body, editedProfile); err != nil {
			log.Fatal(err)
		}

		vars := mux.Vars(r)
		id := vars["id"]
		profile := st.LoadProfile(id)
		if profile == nil {
			fmt.Fprintln(w, "User profile with this ID does not exists!")
			return
		}
		profile.Edit(editedProfile)
	}
}

func remove(st storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		delete(st, id)
	}
}

func getProfiles(st storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		for id := range st {
			profile := st.LoadProfile(id)
			jsonAsBytes, err := json.Marshal(profile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(w, string(jsonAsBytes))
		}
	}
}

func getProfileByID(st storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		json.NewEncoder(w).Encode(st[id])
	}
}
