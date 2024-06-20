package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandler interface {
	RegisterCommand(function func(params Params) error, command string) error
	HandleCommand(command string) error
}

type CommandManager struct {
	bot       *tgbotapi.BotAPI
	callbacks map[string]func(params Params) error
}

func NewCommandManager(bot *tgbotapi.BotAPI) *CommandManager {
	return &CommandManager{bot: bot, callbacks: make(map[string]func(params Params) error)}
}

func (manager *CommandManager) RegisterCommand(function func(params Params) error, command string) {
	manager.callbacks[command] = function
}

func (manager *CommandManager) HandleCommand(update tgbotapi.Update, state *State) error {
	command := update.Message.Command()
	for key, function := range manager.callbacks {
		if command == key {
			return function(Params{nil, &Message{update.Message, manager.bot}, manager.bot, state})
		}
	}
	return nil
}
