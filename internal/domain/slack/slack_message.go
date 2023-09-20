package slack

type SlackMessage struct {
	Channel     string                    `json:"channel"`
	Username    string                    `json:"username"`
	IconUrl     string                    `json:"icon_url"`
	LinkNames   string                    `json:"link_names"`
	Attachments []SlackMessageAttachments `json:"attachments"`
}

type SlackMessageAttachments struct {
	MrkDwnIn []string `json:"mrkdwn_in"`
	Pretext  string   `json:"pretext"`
	Text     string   `json:"text"`
	Color    string   `json:"color"`
}
