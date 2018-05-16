package platform

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"qiniupkg.com/x/errors.v7"
)

const (
	requestTimeout = 30 * time.Second
)

var (
	ErrRespIsNull = errors.New("Response is null") // 这个错误一般是301等
)

type Location struct {
	Ip       string // ip
	Country  string // 国家
	Region   string // 省
	City     string // 市
	District string // 区
	Isp      string // 电信 联通 之类的
}

func (l *Location) ToString() string {
	strs := make([]string, 0)
	appendFunc := func(s string) {
		if len(s) > 0 {
			strs = append(strs, s)
		}
	}
	appendFunc(l.Country)
	appendFunc(l.Region)
	appendFunc(l.City)
	appendFunc(l.District)
	appendFunc(l.Isp)

	return strings.Join(strs, " ")
}

func HttpGet(apiUrl string) (string, error) {
	client := &http.Client{Timeout: requestTimeout}

	resp, err := client.Get(apiUrl)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", ErrRespIsNull
	}
	defer resp.Body.Close()

	bodyBuf := &bytes.Buffer{}

	if _, err = io.Copy(bodyBuf, resp.Body); err != nil {
		return "", err
	}

	return bodyBuf.String(), nil
}

func HttpPost(apiUrl string, form map[string]string) (string, error) {
	client := &http.Client{Timeout: requestTimeout}

	formValues := url.Values{}
	for k, v := range form {
		formValues.Set(k, v)
	}

	resp, err := client.PostForm(apiUrl, formValues)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", ErrRespIsNull
	}
	defer resp.Body.Close()

	bodyBuf := &bytes.Buffer{}

	if _, err = io.Copy(bodyBuf, resp.Body); err != nil {
		return "", err
	}
	return bodyBuf.String(), nil
}
