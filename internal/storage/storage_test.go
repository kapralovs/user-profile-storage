package storage

import (
	"reflect"
	"sync"
	"testing"

	"github.com/kapralovs/user-profile-storage/internal/users"
)

func TestStorage_Load(t *testing.T) {
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
		{},
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
	type fields struct {
		mu sync.Mutex
		Db map[string]*users.Profile
	}
	type args struct {
		p *users.Profile
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Storage{
				mu: tt.fields.mu,
				Db: tt.fields.Db,
			}
			st.Save(tt.args.p)
		})
	}
}

func TestStorage_Edit(t *testing.T) {
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
