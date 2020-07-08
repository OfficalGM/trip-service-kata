package trip

import (
	"errors"
	"reflect"
	"testing"
)

type fakeUserSession struct {
	user *User
	err  error
}

func (sess *fakeUserSession) GetLoggedUser() (*User, error) {
	return sess.user, sess.err
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
		fakeSession           UserSession
		fakeFindTripsByUserFn func(User) ([]Trip, error)
		want                  []Trip
		wantErr               bool
	}{
		{
			name:        "Not logged in user",
			user:        User{},
			fakeSession: &fakeUserSession{nil, nil},
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Logged in user is not a friend of the user",
			user:        User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeSession: &fakeUserSession{&User{name: "River"}, nil},
			want:        []Trip{},
			wantErr:     false,
		},
		{
			name:                  "Logged in user is a friend of the user",
			user:                  User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeSession:           &fakeUserSession{&User{name: "9N"}, nil},
			fakeFindTripsByUserFn: fakeFindTripsByUser([]Trip{{}, {}}, nil),
			want:                  []Trip{{}, {}},
			wantErr:               false,
		},
		{
			name:                  "FindTripsByUser error",
			user:                  User{name: "Eric", friends: []User{{name: "9N"}, {name: "Tzu"}}},
			fakeSession:           &fakeUserSession{&User{name: "9N"}, nil},
			fakeFindTripsByUserFn: fakeFindTripsByUser(nil, errors.New("error")),
			want:                  nil,
			wantErr:               true,
		},
		{
			name:        "Get logged in user error",
			user:        User{},
			fakeSession: &fakeUserSession{nil, errors.New("error")},
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindTripsByUser = tt.fakeFindTripsByUserFn
			got, err := NewService(tt.fakeSession).GetTripsByUser(tt.user)
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
