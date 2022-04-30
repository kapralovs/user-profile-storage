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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAdminRights(tt.args.profile); (err != nil) != tt.wantErr {
				t.Errorf("CheckAdminRights() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
