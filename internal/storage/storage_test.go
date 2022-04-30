package storage

import (
	"reflect"
	"sync"
	"testing"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

func TestStorage_Load(t *testing.T) {
	var (
		strg = New()
	)
	strg.Init()

	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.Profile
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "2",
			},
			want: &users.Profile{
				ID:       "2",
				Email:    "johndoe@domain.com",
				Username: "john_doe",
				Password: "top123secret",
				IsAdmin:  false,
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "dfhjhj",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			got, err := st.Load(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Save(t *testing.T) {
	var (
		strg = New()
	)
	strg.Init()

	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		p *users.Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{&users.Profile{
				ID:       "2",
				Email:    "testUser@domain.com",
				Username: "test_user",
				Password: "superSecretPassword",
				IsAdmin:  false,
			}},
			wantErr: false,
		},
		{
			name: "Empty_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{&users.Profile{
				ID:       "",
				Email:    "johndoe@domain.com",
				Username: "john_doe",
				Password: "top123secret",
				IsAdmin:  true,
			}},
			wantErr: true,
		},
		{
			name: "Nil_Profile_Arg",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: nil,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			if err := st.Save(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Edit(t *testing.T) {
	var (
		strg = New()
	)
	strg.Init()

	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		id string
		np *users.Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "2",
				np: &users.Profile{
					ID:       "2",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  true,
				}},
			wantErr: false,
		},
		{
			name: "Wrong_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "323",
				np: &users.Profile{
					ID:       "323",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  false,
				}},
			wantErr: true,
		},
		{
			name: "Empty_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "",
				np: &users.Profile{
					ID:       "",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  true,
				}},
			wantErr: true,
		},
		{
			name: "Nil_Profile_Arg",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "2",
				np: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			if err := st.Edit(tt.args.id, tt.args.np); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Delete(t *testing.T) {
	var (
		strg = New()
	)
	strg.Init()

	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "2",
			},
			wantErr: false,
		},
		{
			name: "Wrong_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "323",
			},
			wantErr: true,
		},
		{
			name: "Empty_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				id: "",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			if err := st.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_CheckForDuplicates(t *testing.T) {
	var (
		strg = New()
	)
	strg.Init()

	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		p *users.Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: &users.Profile{
					ID:       "4",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "Empty_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				&users.Profile{
					ID:       "",
					Email:    "testUser@domain.com",
					Username: "test_user",
					Password: "superSecretPassword",
					IsAdmin:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "Nil_Profile_Arg",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: nil,
			},
			wantErr: true,
		},
		{
			name: "Duplicate_ID",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: &users.Profile{
					ID: "1",
				},
			},
			wantErr: true,
		},
		{
			name: "Duplicate_Username",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: &users.Profile{
					Username: "john_doe",
				},
			},
			wantErr: true,
		},
		{
			name: "Duplicate_Email",
			fields: fields{
				strg.mu,
				strg.Db,
			},
			args: args{
				p: &users.Profile{
					Email: "mr_robot@domain.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			if err := st.CheckForDuplicates(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.CheckForDuplicates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
