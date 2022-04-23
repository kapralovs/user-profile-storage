package users

type ExampleList []*UserProfile

var someProfiles = ExampleList{
	&UserProfile{
		ID:       "1",
		Email:    "user1@domain.com",
		Username: "user1",
		Password: "password1",
		IsAdmin:  false,
	},
	&UserProfile{
		ID:       "2",
		Email:    "user2@domain.com",
		Username: "user2",
		Password: "password2",
		IsAdmin:  false,
	},
	&UserProfile{
		ID:       "3",
		Email:    "user3@domain.com",
		Username: "user3",
		Password: "password3",
		IsAdmin:  false,
	},
}
