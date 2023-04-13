package notify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/pkg6/go-requests"
	"net/url"
	"strconv"
	"time"
)

//https://open.dingtalk.com/document/robots/custom-robot-access

type DingTalkClient struct {
	AccessToken string
	Secret      string
}
type DingTalkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (t DingTalkClient) I() IClient {
	return &t
}

func (t DingTalkClient) Name() string {
	return "dingtalk"
}

func (t *DingTalkClient) url() *url.URL {
	uri, _ := url.Parse("https://oapi.dingtalk.com/robot/send")
	query := url.Values{}
	query.Set("access_token", t.AccessToken)
	if t.Secret != "" {
		milliTimestamp := time.Now().UnixNano() / 1e6
		stringToSign := fmt.Sprintf("%s\n%s", strconv.Itoa(int(milliTimestamp)), t.Secret)
		mac := hmac.New(sha256.New, []byte(t.Secret))
		mac.Write([]byte(stringToSign))
		sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		query.Set("timestamp", strconv.Itoa(int(milliTimestamp)))
		query.Set("sign", sign)
	}
	uri.RawQuery = query.Encode()
	return uri
}
func (t *DingTalkClient) Send(message IMessage) (result Result, err error) {
	var resp DingTalkResponse
	response, err := requests.PostJson(t.url().String(), message)
	if err != nil {
		result.WithException(err)
		return result, err
	}
	err = response.Unmarshal(&resp)
	result = BuildResult(t.I(), message, resp)
	result.WithException(err)
	return result, err
}
