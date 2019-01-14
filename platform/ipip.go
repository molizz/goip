package platform

import (
	"encoding/json"
	"fmt"
)

type IPIP struct {
	apiURL string
}

func NewIPIP() *IPIP {
	return &IPIP{
		apiURL: "http://freeapi.ipip.net/%s",
	}
}

func (i *IPIP) GetLocation(ip string) (*Location, error) {
	url := fmt.Sprintf(i.apiURL, ip)

	result, err := HttpGet(url)
	if err != nil {
		return nil, err
	}

	var ipInfo []string
	err = json.Unmarshal([]byte(result), &ipInfo)
	if err != nil {
		return nil, err
	}

	if len(ipInfo) < 5 {
		return nil, ErrRespIsError
	}

	location := &Location{
		Ip:      ip,
		Country: ipInfo[0],
		Region:  ipInfo[1],
		City:    ipInfo[2],
		Isp:     ipInfo[4],
	}
	return location, nil
}
