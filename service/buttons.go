package service

import "github.com/moscowHackathon/slack"

// GenerateMessageForSlack2 generates attachement with buttons
func GenerateMessageForSlack2(textMessage, id string) slack.Attachment {
	result := slack.Attachment{
		Text: textMessage,
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "yes",
				Text:  "да",
				Type:  "button",
				Value: "1",
			},
			slack.AttachmentAction{
				Name:  "no",
				Text:  "нет",
				Type:  "button",
				Value: "-1",
			},
			slack.AttachmentAction{
				Name:  "unknown",
				Text:  "хз",
				Type:  "button",
				Value: "0",
			},
		},
		CallbackID: id,
	}
	return result
}
