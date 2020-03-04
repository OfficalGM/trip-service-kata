package trip

import (
	"reflect"
	"testing"
)

func TestService_GetTripsByUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		want    []Trip
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService().GetTripsByUser(tt.user)
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
