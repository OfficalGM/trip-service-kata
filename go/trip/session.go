package trip

import (
	"sync"
)

var (
	sess *Session
	once sync.Once
)

type Session struct{}

func NewSession() *Session {
	once.Do(func() {
		sess = &Session{}
	})
	return sess
}

func (s *Session) IsUserLoggedIn(user User) (bool, error) {
	panic("trip.IsUserLoggedIn should not be called in an unit test")
}

func (s *Session) GetLoggedUser() (*User, error) {
	panic("trip.GetLoggedUser should not be called in an unit test")
}
