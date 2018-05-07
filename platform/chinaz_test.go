package platform

import "testing"

func TestNewChinaz(t *testing.T) {
	chinaz := NewChinaz()
	location, err := chinaz.GetLocation("61.132.63.68")
	if err != nil {
		t.Errorf("%#v", err)
	}
	t.Logf("location %#v", location)
}
