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
	Delete() (*Message, error)
	Answer(params MessageParams) (*Message, error)
	EditText(params MessageParams) (*Message, error)
}

func getMessage(message tgbotapi.Chattable, bot *tgbotapi.BotAPI) (*Message, error) {
	res, err := bot.Send(message)
	if err != nil {
		return nil, err
	} else {
		return NewMessage(&res, bot), err
	}
}

func (message *Message) Answer(params MessageParams) (*Message, error) {
	msg := tgbotapi.NewMessage(message.ApiMessage.Chat.ID, params.Text)
	msg.ReplyMarkup = params.ReplyMarkup
	if params.ParseMode != "" {
		msg.ParseMode = params.ParseMode
	}
	return getMessage(msg, message.bot)
}

func (message *Message) EditText(params MessageParams) (*Message, error) {
	msg := tgbotapi.NewEditMessageText(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID, params.Text)
	msg.ReplyMarkup = params.InlineReplyMarkup
	return getMessage(msg, message.bot)
}

func (message *Message) Delete() (*Message, error) {
	msg := tgbotapi.NewDeleteMessage(message.ApiMessage.Chat.ID, message.ApiMessage.MessageID)
	return getMessage(msg, message.bot)
}
