package update_handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Handler struct {
	Bot             *tgbotapi.BotAPI
	CallbackManager *CallbackManager
	CommandManager  *CommandManager
	states          map[int64]*State
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	handler := &Handler{Bot: bot}
	handler.CallbackManager = NewCallbackManager(bot)
	handler.CommandManager = NewCommandManager(bot)
	handler.states = make(map[int64]*State)
	return handler
}

func (handler *Handler) HandleUpdate(update tgbotapi.Update) error {

	chatId := update.FromChat().ID
	userState, isPresent := handler.states[chatId]

	if !isPresent {
		userState = &State{userId: chatId}
		handler.states[chatId] = userState
	}

	if update.CallbackQuery != nil {
		err := handler.CallbackManager.HandleCallback(update, userState)
		if err != nil {
			return err
		}
	}

	if update.Message != nil {
		err := handler.CommandManager.HandleCommand(update, userState)
		if err != nil {
			return err
		}
	}

	return nil

}
