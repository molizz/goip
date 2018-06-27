package platform

import (
	"encoding/json"
	"fmt"
)

type Tencent struct {
	key    string
	apiUrl string
}

func NewTencent(key string) *Tencent {
	return &Tencent{
		key:    key,
		apiUrl: "http://apis.map.qq.com/ws/location/v1/ip?ip=%s&key=%s",
	}
}

func (t *Tencent) GetLocation(ip string) (*Location, error) {
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
		Ip:       ip,
		Country:  data["nation"].(string),
		Region:   data["province"].(string),
		City:     data["city"].(string),
		District: data["district"].(string),
		//Isp:     data[""],
	}

	if len(location.City) == 0 {
		location.City = location.Region
	}
	if len(location.Region) == 0 {
		location.Region = location.Country
	}

	return location, nil
}

func (t *Tencent) parseToMap(str string) (map[string]interface{}, error) {
	resp := &struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Result  struct {
			AdInfo map[string]interface{} `json:"ad_info"`
		} `json:"result"`
	}{
		Result: struct {
			AdInfo map[string]interface{} `json:"ad_info"`
		}{
			AdInfo: make(map[string]interface{}),
		},
	}

	err := json.Unmarshal([]byte(str), resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, fmt.Errorf("Result code is %d, message %s", resp.Status, resp.Message)
	}

	return resp.Result.AdInfo, nil
}
