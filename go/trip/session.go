package trip

import (
	"errors"
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
	return false, errors.New("trip.IsUserLoggedIn should not be called in an unit test")
}

func (s *Session) GetLoggedUser() (*User, error) {
	return nil, errors.New("trip.GetLoggedUser should not be called in an unit test")
}
