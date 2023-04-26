package wework

import (
	"github.com/pkg6/go-notify"
)

type NewsMessage struct {
	Msgtype string `json:"msgtype"`
	News    struct {
		Articles []Article `json:"articles"`
	} `json:"news"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

func NewNewsMessage(articles []Article) notify.IMessage {
	m := NewsMessage{}
	m.News.Articles = articles
	return m.I()
}

func (t NewsMessage) I() notify.IMessage {
	t.Msgtype = "news"
	return &t
}

func (t *NewsMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (t *NewsMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (t *NewsMessage) AtAll() {
	panic("No delivery required")
}
