package main

import (
	"headquarters/update_handlers"
)

type DeleterInterface interface {
	AddMessage(message *update_handlers.Message)
	DeleteMessages(userId int64)
}

type Deleter struct {
	messages map[int64][]*update_handlers.Message
}

func NewDeleter() *Deleter {
	return &Deleter{messages: make(map[int64][]*update_handlers.Message)}
}

func (d *Deleter) AddMessage(message *update_handlers.Message) {
	userId := message.ApiMessage.Chat.ID
	userMessages, isPresent := d.messages[userId]

	if !isPresent {
		userMessages = make([]*update_handlers.Message, 0)
	}

	userMessages = append(userMessages, message)
	d.messages[userId] = userMessages
}

func (d *Deleter) DeleteMessages(userId int64) {
	for _, msg := range d.messages[userId] {
		_, _ = msg.Delete()
	}
	d.messages[userId] = make([]*update_handlers.Message, 0)
}
