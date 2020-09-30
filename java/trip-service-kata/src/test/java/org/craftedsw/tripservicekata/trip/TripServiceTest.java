package org.craftedsw.tripservicekata.trip;

import org.craftedsw.tripservicekata.exception.CollaboratorCallException;
import org.craftedsw.tripservicekata.user.User;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class TripServiceTest {

    @Test
    public void test_1() {
        TripService tripService = new TripService();

        User user = new User();
        assertThatThrownBy(() -> {
            tripService.getTripsByUser(user);

        }).isInstanceOf(CollaboratorCallException.class);

    }
    
}
