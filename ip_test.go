package goip

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	AddLocal()
	location, err := GetLocation("123.58.180.8")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
}
