package trip_test

import (
	"reflect"
	"testing"

	"tripkata/trip"
)

func TestService_GetTripsByUser(t *testing.T) {
	tests := []struct {
		name    string
		user    trip.User
		want    []trip.Trip
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := trip.NewService().GetTripsByUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTripsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTripsByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
