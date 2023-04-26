package wework

import (
	"github.com/pkg6/go-notify"
)

type TemplateCardTextNoticeMessage struct {
	Msgtype                string                 `json:"msgtype"`
	TemplateCardTextNotice TemplateCardTextNotice `json:"template_card"`
}

type TemplateCardTextNotice struct {
	TemplateCardCommon
	EmphasisContent struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"emphasis_content"`
	SubTitleText string `json:"sub_title_text"`
}

func NewTemplateCardTextNoticeMessage(templateCard TemplateCardTextNotice) notify.IMessage {
	m := TemplateCardTextNoticeMessage{}
	m.TemplateCardTextNotice = templateCard
	return m.I()
}

func (t TemplateCardTextNoticeMessage) I() notify.IMessage {
	t.Msgtype = "template_card"
	return &t
}

func (t *TemplateCardTextNoticeMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (t *TemplateCardTextNoticeMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (t *TemplateCardTextNoticeMessage) AtAll() {
	panic("No delivery required")
}
