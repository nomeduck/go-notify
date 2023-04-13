package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type MarkdownMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At At `json:"at"`
}

func NewMarkdownMessage(title, text string) notify.IMessage {
	m := MarkdownMessage{}
	m.Markdown.Title = title
	m.Markdown.Text = text
	return m.I()
}
func (t MarkdownMessage) I() notify.IMessage {
	t.Msgtype = "markdown"
	return &t
}

func (t *MarkdownMessage) AtMobiles(mobiles []string) {
	t.At.AtMobiles(mobiles)
}

func (t *MarkdownMessage) AtUserIds(userIds []string) {
	t.At.AtUserIds(userIds)
}

func (t *MarkdownMessage) AtAll() {
	t.At.AtAll()
}
