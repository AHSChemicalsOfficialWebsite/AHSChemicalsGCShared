package models

import "encoding/json"

type Notifcation struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (n *Notifcation) ToBytes() []byte {
	b, _ := json.Marshal(n)
	return b
}
