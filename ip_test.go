package goip

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	location, err := AddTaobao().GetLocation("123.58.180.8")
	if err != nil {
		t.Error(err, location.ToString())
	}
}
