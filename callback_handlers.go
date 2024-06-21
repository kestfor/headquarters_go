package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/geo"
	"headquarters/update_handlers"
	"headquarters/utils"
)

func ChallengeCallback(params update_handlers.RedirectedParams) error {
	message := params.Message
	state := params.State
	state.SetState("location")
	state.SetData(map[string]any{"attempts": 0})
	_ = message.Delete()
	return message.Answer(update_handlers.MessageParams{Text: "отправить локацию", ReplyMarkup: &utils.GeolocationReplyKeyboard})
}

func ChooseHouseCallback(params update_handlers.RedirectedParams) error {
	message := params.Message
	params.State.Clear()
	return message.EditText(update_handlers.MessageParams{Text: "выбери штаб", InlineReplyMarkup: &utils.HousesKeyboard})
}

func GoBackCallback(params update_handlers.RedirectedParams) error {
	message := params.Message
	params.State.Clear()
	return message.EditText(update_handlers.MessageParams{Text: "главное меню", InlineReplyMarkup: &utils.MenuKeyboard})
}

func HousesCallback(params update_handlers.RedirectedParams) error {
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

	return params.Message.EditText(update_handlers.MessageParams{Text: "штаб изменен", InlineReplyMarkup: &utils.GoBackKeyboard})
}

func DownloadFiles(params update_handlers.RedirectedParams) error {
	message := params.Message
	bot := params.Bot

	statsFile := tgbotapi.FilePath(DataBase.StatsFileName)
	usersFile := tgbotapi.FilePath(DataBase.UserFileName)
	phrasesFile := tgbotapi.FilePath(DataBase.PhrasesFileName)

	phrases := tgbotapi.NewDocument(message.ApiMessage.Chat.ID, phrasesFile)
	stats := tgbotapi.NewDocument(message.ApiMessage.Chat.ID, statsFile)
	users := tgbotapi.NewDocument(message.ApiMessage.Chat.ID, usersFile)

	_, err1 := bot.Send(stats)
	_, err2 := bot.Send(users)
	_, err3 := bot.Send(phrases)
	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
	}

	if err3 != nil {
		return err3
	}

	return nil
}

func AddPhraseCallback(params update_handlers.RedirectedParams) error {
	message := params.Message
	state := params.State

	state.SetState("phrase")
	return message.EditText(update_handlers.MessageParams{Text: "отправь фразу", InlineReplyMarkup: &utils.GoBackKeyboard})
}
