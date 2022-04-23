package users

type ExampleList []*Profile

var someProfiles = ExampleList{
	&Profile{
		ID:       "1",
		Email:    "user1@domain.com",
		Username: "user1",
		Password: "password1",
		IsAdmin:  false,
	},
	&Profile{
		ID:       "2",
		Email:    "user2@domain.com",
		Username: "user2",
		Password: "password2",
		IsAdmin:  false,
	},
	&Profile{
		ID:       "3",
		Email:    "user3@domain.com",
		Username: "user3",
		Password: "password3",
		IsAdmin:  false,
	},
}
