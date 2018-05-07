package platform

import (
	"strings"
)

type Chinaz struct {
	apiUrl string
}

func NewChinaz() *Chinaz {
	return &Chinaz{
		apiUrl: "http://ip.chinaz.com/ajaxsync.aspx?at=ipbatch",
	}
}

func (c *Chinaz) GetLocation(ip string) (*Location, error) {
	form := map[string]string{
		"ip": ip,
	}

	result, err := HttpPost(c.apiUrl, form)
	if err != nil {
		return nil, err
	}

	lc, err := c.fetchLocation(result)
	if err != nil {
		return nil, err
	}

	location := c.parseRegion(lc)
	location.Ip = ip

	return location, nil
}

func (c *Chinaz) parseRegion(location string) *Location {
	lo := &Location{}

	regionIndex := strings.Index(location, "省")
	if regionIndex >= 0 {
		lo.Region = location[:regionIndex]
	}

	cityIndex := strings.Index(location, "市")
	if cityIndex >= 0 && cityIndex >= regionIndex {
		lo.City = location[regionIndex+3 : cityIndex] // rune = int32 = 3x8bit
	}

	lcs := strings.Split(location, " ")
	if len(lcs) > 0 {
		lo.Isp = lcs[len(lcs)-1]
	}
	return lo
}

func (c *Chinaz) fetchLocation(str string) (location string, err error) {
	str = strings.TrimSuffix(str, "'}])")
	strs := strings.Split(str, "'")
	if l := len(strs); l > 0 {
		return strs[l-1], nil
	}
	return
}
