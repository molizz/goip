package platform

import "testing"

func TestNewTencent(t *testing.T) {
	tencent := NewTencent("XAPBZ-J4DCV")
	location, err := tencent.GetLocation("61.142.163.68")
	if err != nil {
		t.Error(err, location)
	}
}
