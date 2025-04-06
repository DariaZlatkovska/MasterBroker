package model

import "encoding/json"

type BaseMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (m BaseMessage) EventType() string {
	return m.Type
}
