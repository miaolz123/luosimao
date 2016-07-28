package luosimao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	smsURL   = "http://sms-api.luosimao.com/v1/send.json"
	voiceURL = "http://voice-api.luosimao.com/v1/verify.json"
)

// Client ...
type Client struct {
	KeySMS   string
	KeyVoice string
}

// SendSMS ...
func (c Client) SendSMS(mobile interface{}, message string) error {
	params := url.Values{}
	params.Add("mobile", fmt.Sprint(mobile))
	params.Add("message", message)
	body, err := httpDo("POST", smsURL, c.Key, params)
	if err != nil {
		return err
	}
	result := struct {
		Msg   string `json:"msg"`
		Error int    `json:"error"`
	}{"Unkonw Error", -1}
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Error != 0 {
		err = fmt.Errorf(result.Msg)
	}
	return err
}

func httpDo(method, url, key string, params url.Values) (body []byte, err error) {
	c := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", key)
	var resp *http.Response
	resp, err = c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
