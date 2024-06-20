package main

import (
	"headquarters/code/update_handlers"
)

type DeleterInterface interface {
	AddMessage(message update_handlers.Messenger)
	DeleteMessages(userId int64)
}

type Deleter struct {
	messages map[int64][]update_handlers.Messenger
}

func NewDeleter() *Deleter {
	return &Deleter{messages: make(map[int64][]update_handlers.Messenger)}
}

func (d *Deleter) AddMessage(message update_handlers.Messenger) {
	userId := message.GetMessage().Chat.ID
	userMessages, isPresent := d.messages[userId]

	if !isPresent {
		userMessages = make([]update_handlers.Messenger, 0)
	}

	userMessages = append(userMessages, message)
	d.messages[userId] = userMessages
}

func (d *Deleter) DeleteMessages(userId int64) {
	for _, msg := range d.messages[userId] {
		_, _ = msg.Delete()
	}
	d.messages[userId] = make([]update_handlers.Messenger, 0)
}
