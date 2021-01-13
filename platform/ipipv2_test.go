package platform

import (
	"testing"
)

func TestNewIPIPv2(t *testing.T) {
	ipip := NewIPIPv2("")
	loc, err := ipip.GetLocation("119.139.1.1")
	if err != nil {
		t.Fatal(err)
	}
	if loc.Country != "中国" {
		t.Fatal(loc.Country)
	}
	if loc.Region != "广东" {
		t.Fatal(loc.Region)
	}
	if loc.City != "深圳" {
		t.Fatal(loc.City)
	}
}
