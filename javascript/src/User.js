'use strict';

module.exports = class User {
  constructor(trips, friends) {
    this.trips = trips;
    this.friends = friends;
  }

  getFriends() {
    return this.friends;
  }

  addFriend(user) {
    this.friends.push(user);
  }

  addTrip(trip) {
    this.trips.add(trip);
  }

  trips() {
    return this.trips;
  }
};
