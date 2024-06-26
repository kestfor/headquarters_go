package update_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler interface {
	RegisterCallback(function func(params RedirectedParams) error, callbackData string) error
	HandleCallback(callbackData string) error
}

type CallbackFactory struct {
	prefix string
	data   map[any]any
}

type CallbackManager struct {
	bot       *tgbotapi.BotAPI
	callbacks map[string]func(params RedirectedParams) error
}

func NewCallbackManager(bot *tgbotapi.BotAPI) *CallbackManager {
	return &CallbackManager{callbacks: make(map[string]func(params RedirectedParams) error), bot: bot}
}

func (manager *CallbackManager) RegisterCallbackFactory(function func(params RedirectedParams) error, factory *CallbackFactory) {
	manager.callbacks[factory.prefix] = function
}

func (manager *CallbackManager) RegisterCallback(function func(params RedirectedParams) error, callbackData string) {
	manager.callbacks[callbackData] = function
}

func (manager *CallbackManager) HandleCallback(update tgbotapi.Update, state *State) error {
	callbackData := update.CallbackData()
	for key, function := range manager.callbacks {
		if key == callbackData {
			err := function(RedirectedParams{update.CallbackQuery, &Message{update.CallbackQuery.Message, manager.bot}, manager.bot, state})
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
