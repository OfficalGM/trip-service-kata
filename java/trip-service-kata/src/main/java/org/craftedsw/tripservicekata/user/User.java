package org.craftedsw.tripservicekata.user;

import java.util.ArrayList;
import java.util.List;

import org.craftedsw.tripservicekata.trip.Trip;

public class User {

    private List<Trip> trips;
    private List<User> friends;

    public User() {
        trips = new ArrayList<>();
        friends = new ArrayList<User>();
    }

    User(List<Trip> trips, List<User> friends) {
        this.trips = trips;
        this.friends = friends;
    }

    public List<User> getFriends() {
        return friends;
    }

    public void addFriend(User user) {
        friends.add(user);
    }

    public void addTrip(Trip trip) {
        trips.add(trip);
    }

    public List<Trip> trips() {
        return trips;
    }

    public List<Trip> getTrips() {
        return trips;
    }

    public boolean isFriendWith(User friend) {
        return friends.contains(friend);
    }

    public static UserBuilder builder() {
        return new UserBuilder();
    }

    public static class UserBuilder {

        private List<Trip> trips;
        private List<User> friends;

        public UserBuilder trips(List<Trip> tripList) {
            trips = tripList;
            return this;
        }

        public UserBuilder friends(List<User> userList) {
            friends = userList;
            return this;
        }

        public User build() {
            if (trips == null) {
                trips = new ArrayList<>();
            }
            if (friends == null) {
                friends = new ArrayList<>();
            }
            return new User(trips, friends);
        }

    }


}
