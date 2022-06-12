package slack

// Metadata is a structured payload defined by the user that can be sent with messages.
//
// https://api.slack.com/metadata
type Metadata struct {
	Type    string      `json:"event_type"`
	Payload interface{} `json:"event_payload"`
}
