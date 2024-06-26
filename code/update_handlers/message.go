package update_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message struct {
	apiMessage *tgbotapi.Message
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

type Messenger interface {
	GetMessage() *tgbotapi.Message
	Delete() (Messenger, error)
	Answer(params MessageParams) (Messenger, error)
	EditText(params MessageParams) (Messenger, error)
}

func (m *Message) GetMessage() *tgbotapi.Message {
	return m.apiMessage
}

func getMessage(message tgbotapi.Chattable, bot *tgbotapi.BotAPI) (Messenger, error) {
	res, err := bot.Send(message)
	if err != nil {
		return nil, err
	} else {
		return NewMessage(&res, bot), err
	}
}

func (m *Message) Answer(params MessageParams) (Messenger, error) {
	msg := tgbotapi.NewMessage(m.apiMessage.Chat.ID, params.Text)
	msg.ReplyMarkup = params.ReplyMarkup
	if params.ParseMode != "" {
		msg.ParseMode = params.ParseMode
	}
	return getMessage(msg, m.bot)
}

func (m *Message) EditText(params MessageParams) (Messenger, error) {
	msg := tgbotapi.NewEditMessageText(m.apiMessage.Chat.ID, m.apiMessage.MessageID, params.Text)
	msg.ReplyMarkup = params.InlineReplyMarkup
	return getMessage(msg, m.bot)
}

func (m *Message) Delete() (Messenger, error) {
	msg := tgbotapi.NewDeleteMessage(m.apiMessage.Chat.ID, m.apiMessage.MessageID)
	return getMessage(msg, m.bot)
}
