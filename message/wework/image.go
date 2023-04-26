package wework

import (
	"github.com/pkg6/go-notify"
)

type ImageMessage struct {
	Msgtype string `json:"msgtype"`
	Image   struct {
		Base64 string `json:"base64"`
		Md5    string `json:"md5"`
	} `json:"image"`
}

func NewImageMessage(base64, md5 string) notify.IMessage {
	m := ImageMessage{}
	m.Image.Base64 = base64
	m.Image.Md5 = md5
	return m.I()
}

func (t ImageMessage) I() notify.IMessage {
	t.Msgtype = "image"
	return &t
}

func (t *ImageMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (t *ImageMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (t *ImageMessage) AtAll() {
	panic("No delivery required")
}
