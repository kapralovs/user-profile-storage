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
		// TODO: Add test cases.
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
