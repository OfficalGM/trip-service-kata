package org.craftedsw.tripservicekata.user;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.assertj.core.api.Assertions.assertThat;

public class UserTest {

    @Test
    public void no_contain() {
        User user = User.builder().build();

        User paul = new User();

        final boolean result = user.isFriendWith(paul);

        assertThat(result).isFalse();
    }

    @Test
    public void contain() {
        User paul = new User();
        User user = User.builder()
                .friends(Arrays.asList(paul))
                .build();

        final boolean result = user.isFriendWith(paul);

        assertThat(result).isTrue();
    }

}
