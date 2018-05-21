package goip

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	location, err := AddTaobao().GetLocation("35.185.191.24")
	if err != nil {
		t.Error(err, location.ToString())
	}
}
