package platform

import (
	"fmt"

	"github.com/ipipdotnet/datx-go"
)

type Local struct {
	dbPath string
}

func NewLocal() *Local {
	return &Local{
		dbPath: "db/17monipdb.datx",
	}
}

func (c *Local) GetLocation(ip string) (*Location, error) {
	city, err := datx.NewCity(c.dbPath)
	if err != nil {
		return nil, err
	}
	loc, err := city.FindLocation(ip)
	if err != nil {
		return nil, err
	}
	location := &Location{
		Ip:       ip,
		Country:  loc.Country,
		Region:   loc.Province,
		City:     loc.City,
		District: loc.City,
		Isp:      loc.ISP,
	}
	fmt.Println(string(loc.ToJSON()))
	return location, nil
}
