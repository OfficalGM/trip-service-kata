'use strict';

let UserSession = require('./UserSession');
let TripDAO = require('./TripDAO');

class TripService {
  getTripsByUser(user, loggedUser) {
    let tripList = [];
    let isFriend = false;

    if (loggedUser === null) {
      throw new Error('User not logged in.');
    }

    return user.isFriendsWith(loggedUser) ? this._findTripsByUser(user) : [];
  }

  _findTripsByUser() {
    return TripDAO.findTripsByUser();
  }
}

module.exports = TripService;
