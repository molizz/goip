package platform

import (
	"encoding/json"
	"fmt"
)

type Taobao struct {
	apiUrl string
}

func NewTaobao() *Taobao {
	return &Taobao{
		apiUrl: "http://ip.taobao.com/service/getIpInfo.php?ip=%s",
	}
}

func (t *Taobao) GetLocation(ip string) (*Location, error) {
	u := fmt.Sprintf(t.apiUrl, ip)
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
		Country: data["country"],
		Region:  data["region"],
		City:    data["city"],
		Isp:     data["isp"],
	}

	return location, nil
}

func (t *Taobao) parseToMap(str string) (map[string]string, error) {
	resp := &struct {
		Code int               `json:"code"`
		Data map[string]string `json:"data"`
	}{
		Data: make(map[string]string),
	}

	err := json.Unmarshal([]byte(str), resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("Result code is %d", resp.Code)
	}

	return resp.Data, nil
}