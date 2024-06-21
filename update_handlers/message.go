package update_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message struct {
	ApiMessage *tgbotapi.Message
	bot        *tgbotapi.BotAPI
}

type MessageParams struct {
	Text              string
	ReplyMarkup       any
	InlineReplyMarkup *tgbotapi.InlineKeyboardMarkup
	ParseMode         string
}

func NewMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) *Message {
	return &Message{msg, bot}
}

type MessageInterface interface {
	Delete() (tgbotapi.Message, error)
	Answer(params MessageParams) (tgbotapi.Message, error)
	EditText(params MessageParams) (tgbotapi.Message, error)
}

func (message *Message) Answer(params MessageParams) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(message.ApiMessage.Chat.ID, params.Text)
	msg.ReplyMarkup = params.ReplyMarkup
	if params.ParseMode != "" {
		msg.ParseMode = params.ParseMode
	}
	return message.bot.Send(msg)
}

func (message *Message) EditText(params MessageParams) (tgbotapi.Message, error) {
	msg := tgbotapi.NewEditMessageText(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID, params.Text)
	msg.ReplyMarkup = params.InlineReplyMarkup
	return message.bot.Send(msg)
}

func (message *Message) Delete() (tgbotapi.Message, error) {
	msg := tgbotapi.NewDeleteMessage(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID)
	return message.bot.Send(msg)
}
