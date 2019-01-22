package goip

import (
	"errors"
	"time"

	"github.com/molizz/goip/platform"
)

type Platform interface {
	GetLocation(ip string) (*platform.Location, error)
}

type Locator interface {
	GetCountry() string
	GetRegion() string
	GetCity() string

	ToString() string
}

type Address struct {
	plts []Platform
}

var address *Address

func (a *Address) GetLocation(ip string, timeout int) (locator Locator, err error) {
	length := len(address.plts)
	if length == 0 {
		return nil, errors.New("Not found Platform")
	}
	if timeout > 0 {
		platform.Timeout = time.Duration(timeout) * time.Second
	}

	for _, plt := range address.plts {
		locator, err = plt.GetLocation(ip)
		if err != nil {
			continue
		}
		return locator, nil
	}

	return nil, err
}

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

func AddFreeIPIP() *Address {
	addPlatform(platform.NewFreeIPIP())
	return address
}

func AddIPIP(token string) *Address {
	addPlatform(platform.NewIPIP(token))
	return address
}
