package platform

import (
	"fmt"
	"testing"
)

func TestNewTaobao(t *testing.T) {
	tb := NewTaobao()
	loc, err := tb.GetLocation("61.132.63.68")
	if err != nil {
		t.Fatal(err, loc)
	}
	fmt.Println("location: ", loc.Country, loc.Region, loc.City, loc.Isp)
}
