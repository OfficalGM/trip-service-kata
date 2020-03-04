package trip

import (
	"sync"
)

var (
	sess UserSession
	once sync.Once
)

type UserSession interface {
	GetLoggedUser() (*User, error)
}

type Session struct{}

func NewSession() UserSession {
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
