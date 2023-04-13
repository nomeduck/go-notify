package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type TextMessage struct {
	At   At `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func NewTextMessage(content string) notify.IMessage {
	m := TextMessage{}
	m.Text.Content = content
	return m.I()
}

func (t TextMessage) I() notify.IMessage {
	t.Msgtype = "text"
	return &t
}

func (t *TextMessage) AtMobiles(mobiles []string) {
	t.At.AtMobiles(mobiles)
}

func (t *TextMessage) AtUserIds(userIds []string) {
	t.At.AtUserIds(userIds)
}

func (t *TextMessage) AtAll() {
	t.At.AtAll()
}
