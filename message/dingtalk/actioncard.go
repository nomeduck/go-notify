package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type ActionCardMessage struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func NewActionCardMessage(title, text, btnOrientation, SingleTitle, singleURL string) notify.IMessage {
	m := ActionCardMessage{}
	m.ActionCard.Title = title
	m.ActionCard.Text = text
	m.ActionCard.BtnOrientation = btnOrientation
	m.ActionCard.SingleTitle = SingleTitle
	m.ActionCard.SingleURL = singleURL
	return m.I()
}

func (a ActionCardMessage) I() notify.IMessage {
	a.Msgtype = "actionCard"
	return &a
}

func (a *ActionCardMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (a *ActionCardMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (a *ActionCardMessage) AtAll() {
	panic("No delivery required")
}
