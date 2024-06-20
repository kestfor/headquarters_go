package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler interface {
	RegisterCallback(function func(message Message, bot *tgbotapi.BotAPI) error, callbackData string) error
	HandleCallback(callbackData string) error
}

type CallbackManager struct {
	bot       *tgbotapi.BotAPI
	callbacks map[string]func(message Message, bot *tgbotapi.BotAPI) error
}

func NewCallbackManager(bot *tgbotapi.BotAPI) *CallbackManager {
	return &CallbackManager{callbacks: make(map[string]func(message Message, bot *tgbotapi.BotAPI) error), bot: bot}
}

func (manager *CallbackManager) RegisterCallback(function func(message Message, bot *tgbotapi.BotAPI) error, callbackData string) {
	manager.callbacks[callbackData] = function
}

func (manager *CallbackManager) HandleCallback(update tgbotapi.Update) error {
	callbackData := update.CallbackData()
	for key, function := range manager.callbacks {
		if key == callbackData {
			err := function(Message{update.CallbackQuery.Message, manager.bot}, manager.bot)
			if err != nil {
				return err
			}
		}
	}
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, callbackData)
	if _, err := manager.bot.Request(callback); err != nil {
		panic(err)
	}
	return nil
}
