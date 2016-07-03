package service

import "github.com/moscowHackathon/slack"

// GenerateMessageForSlack2 generates attachement with buttons
func GenerateMessageForSlack2(textMessage, id string) slack.Attachment {
	result := slack.Attachment{
		Text:     textMessage,
		ThumbURL: "https://pixabay.com/static/uploads/photo/2016/05/28/05/40/question-mark-1421017_960_720.png",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "yes",
				Text:  "Yes",
				Type:  "button",
				Value: "1",
				Style: "primary",
			},
			slack.AttachmentAction{
				Name:  "no",
				Text:  "No",
				Type:  "button",
				Value: "-1",
				Style: "danger",
			},
			slack.AttachmentAction{
				Name:  "skipped",
				Text:  "Skip",
				Type:  "button",
				Value: "0",
				Style: "default",
			},
		},
		CallbackID: id,
	}
	return result
}
