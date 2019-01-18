package platform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	ipipFreeURL = "http://freeapi.ipip.net/%s"
	ipipURL     = "https://ipapi.ipip.net/find?addr=%s"
)

type ipipResult struct {
	Ret  string   `json:"ret"`
	Data []string `json:"data"`
}

type IPIP struct {
	apiURL string
	token  string // 如果是收费版
	isFree bool   // 免费版
}

func NewFreeIPIP() *IPIP {
	return &IPIP{
		apiURL: ipipFreeURL,
		isFree: true,
	}
}

func NewIPIP(token string) *IPIP {
	return &IPIP{
		apiURL: ipipURL,
		token:  token,
		isFree: false,
	}
}

func (i *IPIP) GetLocation(ip string) (*Location, error) {
	if i.isFree {
		return i.GetFreeLocation(ip)
	} else {
		return i.GetChargedLocation(ip)
	}
}

func (i *IPIP) GetChargedLocation(ip string) (*Location, error) {
	url := fmt.Sprintf(i.apiURL, ip)

	var client http.Client
	client.Timeout = RequestTimeout
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Token", i.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("not found resp")
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	var result ipipResult
	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		return nil, err
	}

	if result.Ret != "ok" || len(result.Data) < 5 {
		return nil, fmt.Errorf("err! ret is %s, ip data is %v", result.Ret, result.Data)
	}

	location := &Location{
		Ip:      ip,
		Country: result.Data[0],
		Region:  result.Data[1],
		City:    result.Data[2],
		Isp:     result.Data[4],
	}
	return location, nil
}

func (i *IPIP) GetFreeLocation(ip string) (*Location, error) {
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
