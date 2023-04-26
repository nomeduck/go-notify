package wework

import (
	"github.com/pkg6/go-notify"
)

type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
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
	t.Text.MentionedMobileList = mobiles
}

func (t *TextMessage) AtUserIds(userIds []string) {
	t.Text.MentionedList = append(t.Text.MentionedList, userIds...)
}

func (t *TextMessage) AtAll() {
	t.Text.MentionedList = append(t.Text.MentionedList, "@all")
}
