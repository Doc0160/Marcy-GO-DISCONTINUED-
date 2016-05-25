package Slack

type (
	OMNI struct {
		Type         string        `json:"type"`
		Subtype      string        `json:"subtype"`
		Channel      string        `json:"channel"`
		User         string        `json:"user"`
		Text         string        `json:"text"`
		TS           *string       `json:"ts"`
		File         *File         `json:"file"`
		Attachements *[]Attachment `json:"attachments"`
		URL          string        `json:"url"`
		ReplyTo      *string       `json:"reply_to"`
		OK           *bool         `json:"ok"`
		Error        *struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"error"`
		Presence *string `json:"presence"`
	}

	Message struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"text"`
		ID      string `json:"id,omitempty"`
	}

	Typing struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		ID      string `json:"id,omitempty"`
	}
)
