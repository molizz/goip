package goip

import (
	"github.com/molizz/goip/platform"
	"testing"
)

func TestGetLocation(t *testing.T) {
	taobeo := platform.NewTaobao()
	platform.NewTencent("key")
	AddPlatform(taobeo)
	location, err := GetLocation("35.185.191.24")
	if err != nil {
		t.Error(err, location)
	}
}
