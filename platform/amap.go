package platform

import (
	"encoding/json"
	"fmt"
)

/**
高德只支持国内ip
*/

type Amap struct {
	key    string
	apiUrl string
}

func NewAmap(key string) *Amap {
	return &Amap{
		key:    key,
		apiUrl: "http://restapi.amap.com/v3/ip?ip=%s&output=json&key=%s",
	}
}

func (t *Amap) GetLocation(ip string) (*Location, error) {
	u := fmt.Sprintf(t.apiUrl, ip, t.key)
	result, err := HttpGet(u)
	if err != nil {
		return nil, err
	}

	data, err := t.parseToMap(result)
	if err != nil {
		return nil, err
	}

	location := &Location{
		Ip:      ip,
		Country: "中国",
		Region:  data["province"].(string),
		City:    data["city"].(string),
	}
	if len(location.City) == 0 {
		location.City = location.Region
	}
	if len(location.Region) == 0 {
		location.Region = location.Country
	}

	return location, nil
}

func (t *Amap) parseToMap(str string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &resp)
	if err != nil {
		return nil, err
	}
	if resp["status"] != "1" {
		return nil, fmt.Errorf("Result code is %d, message %s", resp["status"], resp["info"])
	}

	return resp, nil
}
