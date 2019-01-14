package platform

import (
	"fmt"
	"testing"
)

func TestNewIPIP(t *testing.T) {
	ipip := NewIPIP()
	loc, err := ipip.GetLocation("119.137.53.154")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("location: ", loc.Country, loc.Region, loc.City, loc.Isp)
}
