# A Go (Golang) trip service kata

## Trying to add at least three unit tests for the trip.Service.GetTripsByUser(user) function

- Must return error if can not get user from session
- Return an empty slice of trip.Trip if loggedUser is not a friend of the user
- Return a slice of trip.Trips from DAO if loggedUser is a friend of the user

Feel free to add more tests if you want

## Running unit tests to check the correctness

```shell script
go test trip/* -v
```
