package goip

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	AddLocal()
	location, err := GetLocation("119.137.53.161")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
}
