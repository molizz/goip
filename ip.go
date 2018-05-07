package goip

import (
	"errors"
	"github.com/molizz/goip/platform"
)

type Platform interface {
	GetLocation(ip string) (*platform.Location, error)
}

var plts []Platform

func AddPlatform(obj Platform) {
	plts = append(plts, obj)
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
