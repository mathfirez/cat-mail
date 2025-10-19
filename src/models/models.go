package models

import "time"

type Message struct {
	Content   string
	Author    string
	Posted_on time.Time
}
