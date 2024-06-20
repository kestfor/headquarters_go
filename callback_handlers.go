package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/geo"
	"headquarters/handlers"
	"headquarters/utils"
)

func ChallengeCallback(params handlers.Params) error {
	message := params.Message
	state := params.State
	state.SetState("challenge")
	state.SetData(map[string]any{"attempts": 0})
	_ = message.Delete()
	return message.Answer(handlers.FunctionParams{Text: utils.ChallengeInscription, ReplyMarkup: utils.ChallengeReplyKeyboard, ParseMode: "HTML"})
}

func ChooseHouseCallback(params handlers.Params) error {
	message := params.Message
	return message.EditText(handlers.FunctionParams{Text: "выбери штаб", InlineReplyMarkup: &utils.HousesKeyboard})
}

func GoBackCallback(params handlers.Params) error {
	message := params.Message
	return message.EditText(handlers.FunctionParams{Text: "главное меню", InlineReplyMarkup: &utils.MenuKeyboard})
}

func HousesCallback(params handlers.Params) error {
	callback := params.Callback
	callbackData := callback.Data

	switch callbackData {

	case geo.HomeOfIlya:
		geo.MainHome = geo.Houses[geo.HomeOfIlya]
		break
	case geo.HomeOfAlena:
		geo.MainHome = geo.Houses[geo.HomeOfAlena]
		break
	case geo.HomeOfDima:
		geo.MainHome = geo.Houses[geo.HomeOfDima]
		break
	default:
	}

	return params.Message.EditText(handlers.FunctionParams{Text: "штаб изменен", InlineReplyMarkup: &utils.GoBackKeyboard})
}

func DownloadFiles(params handlers.Params) error {
	message := params.Message
	bot := params.Bot

	statsFile := tgbotapi.FilePath(DataBase.StatsFileName)
	usersFile := tgbotapi.FilePath(DataBase.UserFileName)

	stats := tgbotapi.NewDocument(message.ApiMessage.Chat.ID, statsFile)
	users := tgbotapi.NewDocument(message.ApiMessage.Chat.ID, usersFile)

	_, err1 := bot.Send(stats)
	_, err2 := bot.Send(users)

	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
	}

	return nil
}
