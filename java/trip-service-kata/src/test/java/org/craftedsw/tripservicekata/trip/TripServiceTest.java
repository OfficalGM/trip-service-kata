package org.craftedsw.tripservicekata.trip;

import org.craftedsw.tripservicekata.exception.CollaboratorCallException;
import org.craftedsw.tripservicekata.user.User;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class TripServiceTest {

    User isLoggedUser = new User();

    Trip trip = new Trip();

    @Test
    public void test_1() {
        TripService tripService = new TripService();

        User user = new User();
        assertThatThrownBy(() -> {
            tripService.getTripsByUser(user);

        }).isInstanceOf(CollaboratorCallException.class).hasMessage("UserSession.getLoggedUser() should not be called in an unit test");
    }

    @Test
    public void test_2() {
        TripService tripService = new FakeTripService();

        User user = new User();
        user.addFriend(isLoggedUser);

        assertThatThrownBy(() -> {
            tripService.getTripsByUser(user);

        }).isInstanceOf(CollaboratorCallException.class).hasMessage("TripDAO should not be invoked on an unit test.");
    }

    @Test
    public void test_3() {
        TripService tripService = new FakeTripService2();
        User user = new User();
        user.addFriend(isLoggedUser);
        user.addFriend(new User());

        final List<Trip> resultList = tripService.getTripsByUser(user);

        assertThat(resultList).hasSize(1);
    }

    class FakeTripService extends TripService {

        @Override
        User getLoggedUser() {
            return isLoggedUser;
        }
    }

    class FakeTripService2 extends TripService {
        @Override
        User getLoggedUser() {
            return isLoggedUser;
        }

        @Override
        List<Trip> findTripsByUser(User user) {
            List<Trip> tripList = new ArrayList<>();
            tripList.add(trip);
            return tripList;
        }
    }

}
