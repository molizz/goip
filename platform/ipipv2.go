package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	ipipV2URL = "https://ipapi.ipip.net/v2/query/%s?token=%s"
)

type ipipV2Result struct {
	Ret  string `json:"ret"`
	Data struct {
		IP   string `json:"ip"`
		Info struct {
			CountryName string `json:"country_name"`
			RegionName  string `json:"region_name"`
			CityName    string `json:"city_name"`
			Line        string `json:"line"`
		} `json:"info"`
	} `json:"data"`
}

type IPIPv2 struct {
	apiURL string
	token  string
}

func NewIPIPv2(token string) *IPIPv2 {
	return &IPIPv2{
		apiURL: ipipV2URL,
		token:  token,
	}
}

func (i *IPIPv2) GetLocation(ip string) (*Location, error) {
	return i.GetChargedLocation(ip)
}

func (i *IPIPv2) GetChargedLocation(ip string) (*Location, error) {
	url := fmt.Sprintf(i.apiURL, ip, i.token)

	var client http.Client
	client.Timeout = Timeout
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("not found resp")
	}
	defer resp.Body.Close()

	var result ipipV2Result

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Ret != "ok" {
		return nil, fmt.Errorf("err! ret is %s, ip data is %v", result.Ret, result.Data)
	}

	location := &Location{
		Ip:      ip,
		Country: result.Data.Info.CountryName,
		Region:  result.Data.Info.RegionName,
		City:    result.Data.Info.CityName,
		Isp:     result.Data.Info.Line,
	}
	return location, nil
}
