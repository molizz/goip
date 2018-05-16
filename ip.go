package goip

import (
	"errors"
	"github.com/molizz/goip/platform"
)

type Platform interface {
	GetLocation(ip string) (*platform.Location, error)
}

type Address struct {
	plts []Platform
}

var address *Address

func init() {
	address = &Address{
		plts: make([]Platform, 0, 1),
	}
}

func addPlatform(p Platform) {
	address.plts = append(address.plts, p)
}

func AddTaobao() *Address {
	addPlatform(platform.NewTaobao())
	return address
}

func AddAmap(key string) *Address {
	addPlatform(platform.NewAmap(key))
	return address
}

func AddTencent(key string) *Address {
	addPlatform(platform.NewTencent(key))
	return address
}

func AddChinaz() *Address {
	addPlatform(platform.NewChinaz())
	return address
}

func (a *Address) GetLocation(ip string) (location *platform.Location, err error) {
	length := len(a.plts)
	if length == 0 {
		return nil, errors.New("Not found Platform")
	}

	for _, plt := range a.plts {
		location, err = plt.GetLocation(ip)
		if err != nil {
			continue
		}
		return location, nil
	}

	return nil, err
}
