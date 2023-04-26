package notify

import (
	"errors"
	"github.com/pkg6/go-requests"
	"net/url"
)

const Wework = "wework"

// https://developer.work.weixin.qq.com/document/path/91770

type WeWorkClient struct {
	Key string
}
type WeWorkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (c WeWorkClient) I() IClient {
	return &c
}

func (c WeWorkClient) Name() string {
	return Wework
}

func (c *WeWorkClient) url() *url.URL {
	uri, _ := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/webhook/send")
	query := url.Values{}
	query.Set("key", c.Key)
	uri.RawQuery = query.Encode()
	return uri
}
func (c *WeWorkClient) Send(message IMessage) (result IResult) {
	var resp WeWorkResponse
	result = BuildResult(c.I(), message)
	response, err := requests.PostJson(c.url().String(), message)
	if err != nil {
		result.WithException(err)
		return result
	}
	err = response.Unmarshal(&resp)
	if err != nil {
		result.WithException(err)
		return result
	}
	result.WithResult(resp)
	if resp.Errcode != 0 {
		result.WithException(errors.New(resp.Errmsg))
	}
	return result
}
