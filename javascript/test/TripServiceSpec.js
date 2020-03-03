'use strict';

let assert = require('assert');
let TripService = require('../src/TripService');
const Trip = require('../src/Trip');
const User = require('../src/User');
const NOT_LOGGED_UESR = null;
const LOGGED_USER = new User('John');
let loggedUser = new User('John');
let targetUser = new User('Mery');

class TestTripService extends TripService {
  _getLoggedUser() {
    return loggedUser;
  }
  _findTripsByUser(user) {
    return user.trips;
  }
}
let tripService = new TestTripService();

describe('TripService', () => {
  it('should... ', () => {
    assert.equal(2 + 2, 4);
  });

  it('should trow exception when user is not logged in', () => {
    loggedUser = NOT_LOGGED_UESR;
    assert.throws(() => {
      tripService.getTripsByUser(targetUser);
    });
  });

  it('should not return trips when logged User is not a friend', () => {
    loggedUser = LOGGED_USER; // without friends
    const trips = tripService.getTripsByUser(targetUser);
    assert.equal(trips.length, 0);
  });

  it('should return trips when logged User is a friend', () => {
    loggedUser = LOGGED_USER; // without friends
    const targetUser = new User('Mery', [new Trip(), new Trip()], [loggedUser]);
    loggedUser.addFriend(targetUser);

    const trips = tripService.getTripsByUser(targetUser);

    assert.equal(trips.length, 2);
  });
});
