package goip

import (
	"testing"
	"fmt"
)

func TestGetLocation(t *testing.T) {
	AddTaobao()
	location, err := GetLocation("123.58.180.8")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
}
