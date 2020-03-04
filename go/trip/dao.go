package trip

import (
	"errors"
)

func FindTripsByUser(user User) ([]Trip, error) {
	return nil, errors.New("trip.FindTripsByUser should not be invoked on an unit test")
}
