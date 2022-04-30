package users

import "testing"

func TestCheckAdminRights(t *testing.T) {
	type args struct {
		profile *Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Is_Admin",
			args: args{
				profile: &Profile{
					ID:       "2",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "Is_Not_Admin",
			args: args{
				profile: &Profile{
					ID:       "2",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAdminRights(tt.args.profile); (err != nil) != tt.wantErr {
				t.Errorf("CheckAdminRights() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
