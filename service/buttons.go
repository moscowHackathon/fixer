package service

import "github.com/nlopes/slack"

func GenerateMessageForSlack(textMessage string) slack.Attachment {
	/*
	{
	    "text": "Would you like to play a game?",
	    "attachments": [
		{
		    "text": "Choose a game to play",
		    "fallback": "You are unable to choose a game",
		    "callback_id": "wopr_game",
		    "color": "#3AA3E3",
		    "attachment_type": "default",
		    "actions": [
			{
			    "name": "chess",
			    "text": "Chess",
			    "type": "button",
			    "value": "chess"
			},
			{
			    "name": "maze",
			    "text": "Falken's Maze",
			    "type": "button",
			    "value": "maze"
			},
			{
			    "name": "war",
			    "text": "Thermonuclear War",
			    "style": "danger",
			    "type": "button",
			    "value": "war",
			    "confirm": {
				"title": "Are you sure?",
				"text": "Wouldn't you prefer a good game of chess?",
				"ok_text": "Yes",
				"dismiss_text": "No"
			    }
			}
		    ]
		}
	    ]
	}
	*/

	result := slack.Attachment{
		Text: "Choose a game to play",
		Fallback: "You are unable to choose a game",
		Callback_id: "wopr_game",
		Color: "#3AA3E3",
		Attachment_type: "default",
		Actions: []slack.Action{
			{
				Name: "chess",
				Text: "Chess",
				Type: "button",
				Value: "chess",
			},
			{
				Name: "maze",
				Text: "Falken's Maze",
				Type: "button",
				Value: "maze",
			},
			{
				Name: "war",
				Text: "Thermonuclear War",
				Style: "danger",
				Type: "button",
				Value: "war",
				Confirms: slack.Confirm{
					Title: "Are you sure?",
					Text: "Wouldn't you prefer a good game of chess?",
					Ok_text: "Yes",
					Dismiss_text: "No",
				},
			},
		},

	}

	return result

}
//
//type AttachmentField struct {
//	Title string `json:"title"`
//	Value string `json:"value"`
//	Short bool   `json:"short"`
//}
//
//type Attachment struct {
//	Color    string `json:"color,omitempty"`
//	Fallback string `json:"fallback"`
//
//	AuthorName    string `json:"author_name,omitempty"`
//	AuthorSubname string `json:"author_subname,omitempty"`
//	AuthorLink    string `json:"author_link,omitempty"`
//	AuthorIcon    string `json:"author_icon,omitempty"`
//
//	Title     string `json:"title,omitempty"`
//	TitleLink string `json:"title_link,omitempty"`
//	Pretext   string `json:"pretext,omitempty"`
//	Text      string `json:"text"`
//
//	ImageURL string `json:"image_url,omitempty"`
//	ThumbURL string `json:"thumb_url,omitempty"`
//
//	Fields     []AttachmentField `json:"fields,omitempty"`
//	MarkdownIn []string          `json:"mrkdwn_in,omitempty"`
//}
//



type MyAttachment struct {
	Text string `json:"text,omitempty"`
	Fallback string `json:"fallback,omitempty"`
	Callback_id string `json:"callback_id,omitempty"`
	Color string `json:"color,omitempty"`
	Attachment_type string `json:"attachment_type,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Action struct {
	Name string `json:"name,omitempty"`
	Text string `json:"text,omitempty"`
	Style string `json:"style,omitempty"`
	Type string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Confirms []Confirm `json:"confirm,omitempty"`
}

type Confirm struct{
	Title string `json:"title,omitempty"`
	Text string `json:"text,omitempty"`
	Ok_text string `json:"ok_text,omitempty"`
	Dismiss_text string `json:"dismiss_text,omitempty"`
}

/*
// Attachment contains all the information for an attachment
type Attachment struct {
	Color    string `json:"color,omitempty"`
	Fallback string `json:"fallback"`

	AuthorName    string `json:"author_name,omitempty"`
	AuthorSubname string `json:"author_subname,omitempty"`
	AuthorLink    string `json:"author_link,omitempty"`
	AuthorIcon    string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`
	Pretext   string `json:"pretext,omitempty"`
	Text      string `json:"text"`

	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`

	Fields     []AttachmentField `json:"fields,omitempty"`
	MarkdownIn []string          `json:"mrkdwn_in,omitempty"`
}

 */
