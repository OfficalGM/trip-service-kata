package trip

type User struct {
	name    string
	friends []User
	trips   []Trip
}

func NewUser(name string, friends []User, trips []Trip) *User {
	return &User{name: name, friends: friends, trips: trips}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetFriends() []User {
	return u.friends
}

func (u *User) GetTrips() []Trip {
	return u.trips
}
