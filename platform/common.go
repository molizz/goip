package platform

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"qiniupkg.com/x/errors.v7"
	"time"
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
