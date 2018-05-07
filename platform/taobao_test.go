package platform

import "testing"

func TestNewTaobao(t *testing.T) {
	tb := NewTaobao()
	lc, err := tb.GetLocation("61.132.63.68")
	if err != nil {
		t.Fatal(err, lc)
	}
}
