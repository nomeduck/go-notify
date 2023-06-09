<h1 align="center">Go Notify</h1>

<p align="center">:calling: 一款满足你的多种渠道消息通知</p>


[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/go-notify)](https://goreportcard.com/report/github.com/pkg6/go-notify)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/go-notify?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/go-notify/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/go-notify?badge)
[![Release](https://img.shields.io/github/release/pkg6/go-notify.svg?style=flat-square)](https://github.com/pkg6/go-notify/releases)


## 安装

```
$ go get github.com/pkg6/go-notify
```

## 使用



## 平台支持

* [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)



## 使用
<details>
<summary><b>钉钉群机器人</b></summary>

```
package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/message/dingtalk"
)

func main() {
	client := &notify.DingTalkClient{
		AccessToken: "27bbe68cc8b57acc2973b59fd7ae2460fb0b2322ce2e8660f5fb5b75aee04e88",
		Secret:      "SEC55f77c19089ef4aee0be143a77d12730f2daaa2390b212cffb1e1ac1f23f8ccc",
	}
	message := &dingtalk.TextMessage{}
	message.Text.Content = "测试发送dingtalk2"

	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
```
</details>


## 自定义客户端

~~~
type IClient interface {
	I() IClient
	Name() string
	Send(message IMessage) IResult
}
~~~

## 自定义消息内容

~~~
// IMessage 消息
type IMessage interface {
	// I 初始化并返回IMessage
	I() IMessage
	// AtMobiles 消息需要@人员手机号
	AtMobiles(mobiles []string)
	// AtUserIds 消息需要@人员userid
	AtUserIds(userIds []string)
	// AtAll @所有人
	AtAll()
}
~~~
