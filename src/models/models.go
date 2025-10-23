package models

import "time"

type Message struct {
	Id        string
	Content   string
	Author    string
	Receiver  string
	PostedOn  time.Time
	Printed   bool
	PrintedOn time.Time
}

// Message that goes back to the client. Used so we dont send unused data in the json.
type ClientMessage struct {
	Id       string
	Content  string
	Author   string
	PostedOn time.Time
}
