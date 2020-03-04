package trip

import (
	"errors"
	"reflect"
	"testing"
)

func fakeGetLoggedUser(user *User, err error) func() (*User, error) {
	return func() (*User, error) {
		return user, err
	}
}

func fakeFindTripsByUser(trips []Trip, err error) func(User) ([]Trip, error) {
	return func(User) ([]Trip, error) {
		return trips, err
	}
}

func TestService_GetTripsByUser(t *testing.T) {

	tests := []struct {
		name                  string
		user                  User
		fakeGetLoggedUserFn   func() (*User, error)
		fakeFindTripsByUserFn func(User) ([]Trip, error)
		want                  []Trip
		wantErr               bool
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
		{
			name:                  "Logged in user is a friend of the user",
			user:                  User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeGetLoggedUserFn:   fakeGetLoggedUser(&User{name: "9N"}, nil),
			fakeFindTripsByUserFn: fakeFindTripsByUser([]Trip{{}, {}}, nil),
			want:                  []Trip{{}, {}},
			wantErr:               false,
		},
		{
			name:                  "FindTripsByUser error",
			user:                  User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeGetLoggedUserFn:   fakeGetLoggedUser(&User{name: "9N"}, nil),
			fakeFindTripsByUserFn: fakeFindTripsByUser(nil, errors.New("error")),
			want:                  nil,
			wantErr:               true,
		},
		{
			name:                "Get logged in user error",
			user:                User{},
			fakeGetLoggedUserFn: fakeGetLoggedUser(nil, errors.New("error")),
			want:                nil,
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getLoggedUser = tt.fakeGetLoggedUserFn
			FindTripsByUser = tt.fakeFindTripsByUserFn
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
