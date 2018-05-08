package goip

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	AddTaobao()
	location, err := GetLocation("35.185.191.24")
	if err != nil {
		t.Error(err, location)
	}
}
