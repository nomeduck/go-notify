package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/message/dingtalk"
)

func dingTalk() {
	client := notify.DingTalkClient{
		AccessToken: "06616a181687b5acf2dda36b1ce383c7bea2768da4ff56fc7dbedd62ea9b153b",
		Secret:      "SEC1872aac880182595a5c2614cfef500d97e03898bb296bc28a4050ea0a3153843",
	}.I()
	message := dingtalk.TextMessage{}
	message.Text.Content = "测试发送dingtalk"

	n := notify.New(client)
	sender, _ := n.Sender(message.I())
	for name, result := range sender {
		fmt.Println(name)
		if response, ok := result.Result.Response.(notify.DingTalkResponse); ok {
			fmt.Println(response)
		}
	}
}

func main() {
	dingTalk()
}
