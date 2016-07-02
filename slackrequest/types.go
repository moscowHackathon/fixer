package slackrequest

import "github.com/101nl/slack"

// Payload Action URL invocation payload
type Payload struct {
	Actions         []ActionHash  `json:"actions"`
	CallbackID      string        `json:"callback_id"`
	Team            TeamHash      `json:"team"`
	Channel         ChannelHash   `json:"channel"`
	User            UserHash      `json:"user"`
	ActionTS        string        `json:"action_ts"`
	MessageTS       string        `json:"message_ts"`
	AttachmentID    string        `json:"attachment_id"`
	Token           string        `json:"token"`
	OriginalMessage slack.Message `json:"original_message"`
	ResponseURL     string        `json:"response_url"`
}

// ActionHash action that were clicked
type ActionHash struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TeamHash A small set of string attributes about the team where this action occurred
type TeamHash struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

// ChannelHash Where it all happened — the user inciting this action clicked a button on a message contained within
// a channel, and this hash presents attributed about that channel.
type ChannelHash struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserHash The clicker! The action-invoker! The button-presser!
type UserHash struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
