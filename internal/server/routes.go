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

func create(st *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		newUser := &users.Profile{}
		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Fatal(err)
		}

		if err := st.CheckForDuplicates(newUser); err != nil {
			fmt.Fprintln(w, err)
			return
		}

		log.Printf("New profile created by user \"%s\"", user.Username)
		st.Save(newUser)
		fmt.Fprintln(w, "User profile is created!")
	}
}

func edit(st *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			fmt.Fprintln(w, err)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		editedProfile := &users.Profile{}
		if err := json.Unmarshal(body, editedProfile); err != nil {
			log.Fatal(err)
		}

		vars := mux.Vars(r)
		id := vars["id"]

		if err := st.Edit(id, editedProfile); err != nil {
			fmt.Println(w, err)
		}

		fmt.Fprintln(w, "User profile edited!")
	}
}

func remove(st *storage.Storage) func(http.ResponseWriter, *http.Request) {
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

		if err := st.Delete(id); err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, "User profile is deleted!")
	}
}

func getProfiles(st *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		for id := range st.Db {
			profile, err := st.Load(id)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			jsonAsBytes, err := json.Marshal(profile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(w, string(jsonAsBytes))
		}
	}
}

func getProfileByID(st *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authorization(st, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		profile, err := st.Load(id)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		profileAsBytes, err := json.Marshal(profile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(profileAsBytes))
	}
}
