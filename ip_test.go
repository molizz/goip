package goip

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	location, err := AddFreeIPIP().GetLocation("123.58.180.8", 60)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
}

func TestTencentLocation(t *testing.T) {
	location, err := AddTencent("Y2ABZ-ENY3I-UBNG7-5FNFZ-SZENZ-2ZFTX").GetLocation("123.58.180.8", 60)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("国: %s 省: %s 城: %s", location.GetCountry(), location.GetRegion(), location.GetCity())
	fmt.Println(location.ToString())
}

func TestGetLocationFreeIPIP(t *testing.T) {
	location, err := AddIPIP("token").GetLocation("123.58.180.8", 60)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
	fmt.Println(location.GetCity())
}

func TestGetLocationFreeIPIPv2(t *testing.T) {
	location, err := AddIPIPv2("token").GetLocation("119.139.1.1", 10)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(location.ToString())
	fmt.Println(location.GetCity())
}
