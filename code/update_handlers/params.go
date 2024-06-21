package update_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RedirectedParams struct {
	Callback *tgbotapi.CallbackQuery
	Message  *Message
	Bot      *tgbotapi.BotAPI
	State    *State
}
