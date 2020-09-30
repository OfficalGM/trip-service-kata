package org.craftedsw.tripservicekata.trip;

import org.craftedsw.tripservicekata.exception.CollaboratorCallException;
import org.craftedsw.tripservicekata.user.User;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class TripServiceTest {

    User isLoggedUser = new User();

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


    class FakeTripService extends TripService {

        @Override
        User getLoggedUser() {
            return isLoggedUser;
        }
    }


}
