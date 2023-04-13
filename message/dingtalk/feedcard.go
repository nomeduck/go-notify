package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type Link struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}
type FeedCardBtnMessage struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []Link `json:"links"`
	} `json:"feedCard"`
}

func NewFeedCardBtnMessage(link []Link) notify.IMessage {
	m := FeedCardBtnMessage{}
	m.FeedCard.Links = link
	return m.I()
}

func (a *FeedCardBtnMessage) AddLink(link Link) {
	a.FeedCard.Links = append(a.FeedCard.Links, link)
}
func (a FeedCardBtnMessage) I() notify.IMessage {
	a.Msgtype = "feedCard"
	return &a
}

func (a *FeedCardBtnMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (a *FeedCardBtnMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (a *FeedCardBtnMessage) AtAll() {
	panic("No delivery required")
}
