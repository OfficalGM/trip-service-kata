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
		{
			name:                "Logged in user is not a friend of the user",
			user:                User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeGetLoggedUserFn: fakeGetLoggedUser(&User{name: "River"}, nil),
			want:                []Trip{},
			wantErr:             false,
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
