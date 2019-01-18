package goip

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {

	location, err := AddFreeIPIP().GetLocation("123.58.180.8")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
}

func TestGetLocationForIPIP(t *testing.T) {
	location, err := AddIPIP("token").GetLocation("123.58.180.8")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
	fmt.Println(location.GetCity())
}
