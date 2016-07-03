package service

import "github.com/moscowHackathon/slack"

func GenerateMessageForSlack2(textMessage string) slack.Attachment {
	result := slack.Attachment{
		Text: textMessage,
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "yes",
				Text:  "да",
				Type:  "button",
				Value: "yes",
			},
			slack.AttachmentAction{
				Name:  "no",
				Text:  "нет",
				Type:  "button",
				Value: "maze",
			},
			slack.AttachmentAction{
				Name:  "unknown",
				Text:  "хз",
				Type:  "button",
				Value: "unknown",
			},
		},
		CallbackID: "123",
	}
	return result
}

func GenerateMessageForSlack(textMessage string) slack.Attachment {
	result := slack.Attachment{
		Text:       "Choose a game to play",
		Fallback:   "You are unable to choose a game",
		CallbackID: "wopr_game",
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			{
				Name:  "chess",
				Text:  "Chess",
				Type:  "button",
				Value: "chess",
			},
			{
				Name:  "maze",
				Text:  "Falken's Maze",
				Type:  "button",
				Value: "maze",
			},
			{
				Name:  "war",
				Text:  "Thermonuclear War",
				Style: "danger",
				Type:  "button",
				Value: "war",
				Confirm: slack.AttachmentActionConfirm{
					Title:       "Are you sure?",
					Text:        "Wouldn't you prefer a good game of chess?",
					OkText:      "Yes",
					DismissText: "No",
				},
			},
		},
	}

	return result

}
