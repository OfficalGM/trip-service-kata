package org.craftedsw.tripservicekata.trip;

import org.craftedsw.tripservicekata.exception.UserNotLoggedInException;
import org.craftedsw.tripservicekata.user.User;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class TripServiceTest {
    private static final User UNUSED_USER = null;
    private static final User NON_LOGGED_USER = null;
    private static User loggedUser = new User();
    private static User targetUser = new User();
    private static TripService tripService;


    @BeforeAll
    public static void beforeAllTests() {
        tripService = createTripService();
    }

    @Test
    public void
    shouldThrowExceptionWhenUserIsNotLoggedIn() {
        loggedUser = NON_LOGGED_USER;
        UserNotLoggedInException thrown = assertThrows(
                UserNotLoggedInException.class,
                () -> tripService.getTripsByUser(UNUSED_USER)
        );

    }

    @Test public void
    shouldNotReturnTripsWhenLoggedUserIsNotAFriend() throws Exception {
        List<Trip> trips = tripService.getTripsByUser(targetUser);

        assertEquals(trips.size(), 0);
    }

    @Test
    public void onePlusOneShouldBeTwo() {
        assertEquals(2, 1 + 1);
    }

    @Test
    public void onePlusTwoShouldBeThree() {
        assertEquals(3, 1 + 2);
    }

    private static TripService createTripService() {
        return new TripService() {
            @Override
            protected User loggedUser() {
                return loggedUser;
            }

            @Override
            protected List<Trip> findTripsByUser(User user) {
                return user.trips();
            }
        };
    }
}
