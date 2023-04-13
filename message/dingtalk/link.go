package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type LinkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicURL     string `json:"picUrl"`
		MessageURL string `json:"messageUrl"`
	} `json:"link"`
}

func NewLinkMessage(text, title, picURL, messageURL string) notify.IMessage {
	m := LinkMessage{}
	m.Link.Text = text
	m.Link.Title = title
	m.Link.PicURL = picURL
	m.Link.MessageURL = messageURL
	return m.I()
}

func (l LinkMessage) I() notify.IMessage {
	l.Msgtype = "link"
	return &l
}

func (l *LinkMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (l *LinkMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (l *LinkMessage) AtAll() {
	panic("No delivery required")
}
