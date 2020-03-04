package trip

import (
	"errors"
)

var getLoggedUser = func() (*User, error) {
	return NewSession().GetLoggedUser()
}

type Trip struct{}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTripsByUser(user User) ([]Trip, error) {
	loggedUser, err := getLoggedUser()
	if err != nil {
		return nil, err
	}

	if loggedUser == nil {
		return nil, errors.New("loggedUser is nil")
	}

	isFriend := false
	for _, friend := range user.GetFriends() {
		if friend.GetName() == loggedUser.GetName() {
			isFriend = true
			break
		}
	}

	trips := []Trip{}
	if isFriend {
		trips, err = FindTripsByUser(user)
		if err != nil {
			return nil, err
		}
	}

	return trips, nil
}
