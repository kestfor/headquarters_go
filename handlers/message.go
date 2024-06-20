package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message struct {
	ApiMessage *tgbotapi.Message
	bot        *tgbotapi.BotAPI
}

type FunctionParams struct {
	Text              string
	ReplyMarkup       any
	InlineReplyMarkup *tgbotapi.InlineKeyboardMarkup
	ParseMode         string
}

type MessageInterface interface {
	Delete() error
	Answer(params FunctionParams) error
	EditText(params FunctionParams) error
}

func (message *Message) Answer(params FunctionParams) error {
	msg := tgbotapi.NewMessage(message.ApiMessage.Chat.ID, params.Text)
	msg.ReplyMarkup = params.ReplyMarkup
	if params.ParseMode != "" {
		msg.ParseMode = params.ParseMode
	}
	_, err := message.bot.Send(msg)
	return err
}

func (message *Message) EditText(params FunctionParams) error {
	msg := tgbotapi.NewEditMessageText(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID, params.Text)
	msg.ReplyMarkup = params.InlineReplyMarkup
	_, err := message.bot.Send(msg)
	return err
}

func (message *Message) Delete() error {
	msg := tgbotapi.NewDeleteMessage(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID)
	_, err := message.bot.Send(msg)
	return err
}
