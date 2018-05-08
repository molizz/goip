package goip

import (
	"errors"
	"github.com/molizz/goip/platform"
)

type Platform interface {
	GetLocation(ip string) (*platform.Location, error)
}

var plts []Platform

func addPlatform(p Platform) {
	plts = append(plts, p)
}

func AddTaobao() {
	addPlatform(platform.NewTaobao())
}

func AddAmap(key string) {
	addPlatform(platform.NewAmap(key))
}

func AddTencent(key string) {
	addPlatform(platform.NewTencent(key))
}

func AddChinaz() {
	addPlatform(platform.NewChinaz())
}

func GetLocation(ip string) (location *platform.Location, err error) {
	length := len(plts)
	if length == 0 {
		return nil, errors.New("Not found Platform")
	}

	for _, plt := range plts {
		location, err = plt.GetLocation(ip)
		if err != nil {
			continue
		}
		return location, nil
	}

	return nil, err
}
