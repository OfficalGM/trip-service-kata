package trip

import (
	"reflect"
	"testing"
)

func fakeGetLoggedUser(user *User, err error) func() (*User, error) {
	return func() (*User, error) {
		return user, err
	}
}

func TestService_GetTripsByUser(t *testing.T) {

	tests := []struct {
		name                string
		user                User
		fakeGetLoggedUserFn func() (*User, error)
		want                []Trip
		wantErr             bool
	}{
		{
			name:                "Not logged in user",
			user:                User{},
			fakeGetLoggedUserFn: fakeGetLoggedUser(nil, nil),
			want:                nil,
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getLoggedUser = tt.fakeGetLoggedUserFn
			got, err := NewService().GetTripsByUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTripsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTripsByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
