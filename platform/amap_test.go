package platform

import "testing"

func TestNewAmap(t *testing.T) {
	amap := NewAmap("xxxx")
	location, err := amap.GetLocation("61.232.163.68")
	if err != nil {
		t.Error(err, location)
	}
}
