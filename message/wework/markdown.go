package wework

import (
	"github.com/pkg6/go-notify"
)

type MarkdownMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

func NewMarkdownMessage(content string) notify.IMessage {
	m := MarkdownMessage{}
	m.Markdown.Content = content
	return m.I()
}

func (t MarkdownMessage) I() notify.IMessage {
	t.Msgtype = "markdown"
	return &t
}

func (t *MarkdownMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (t *MarkdownMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (t *MarkdownMessage) AtAll() {
	panic("No delivery required")
}
