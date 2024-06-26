package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/code/geo"
	update_handlers2 "headquarters/code/update_handlers"
	"headquarters/code/utils"
)

func ChallengeCallback(params update_handlers2.RedirectedParams) error {
	message := params.Message
	state := params.State
	state.SetState("location")
	state.SetData(map[string]any{"attempts": 0})
	_, _ = message.Delete()
	msg, err := message.Answer(update_handlers2.MessageParams{Text: "отправить локацию", ReplyMarkup: &utils.GeolocationReplyKeyboard})
	MessageDeleter.AddMessage(msg)
	return err
}

func ChooseHouseCallback(params update_handlers2.RedirectedParams) error {
	message := params.Message
	params.State.Clear()
	_, err := message.EditText(update_handlers2.MessageParams{Text: "выбери штаб", InlineReplyMarkup: &utils.HousesKeyboard})
	return err
}

func GoBackCallback(params update_handlers2.RedirectedParams) error {
	message := params.Message
	params.State.Clear()
	_, err := message.EditText(update_handlers2.MessageParams{Text: "главное меню", InlineReplyMarkup: &utils.MenuKeyboard})
	return err
}

func HousesCallback(params update_handlers2.RedirectedParams) error {
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

	_, err := params.Message.EditText(update_handlers2.MessageParams{Text: "штаб изменен", InlineReplyMarkup: &utils.GoBackKeyboard})
	return err
}

func DownloadFiles(params update_handlers2.RedirectedParams) error {
	message := params.Message
	bot := params.Bot

	statsFile := tgbotapi.FilePath(DataBase.StatsFileName)
	usersFile := tgbotapi.FilePath(DataBase.UserFileName)
	phrasesFile := tgbotapi.FilePath(DataBase.PhrasesFileName)

	phrases := tgbotapi.NewDocument(message.GetMessage().Chat.ID, phrasesFile)
	stats := tgbotapi.NewDocument(message.GetMessage().Chat.ID, statsFile)
	users := tgbotapi.NewDocument(message.GetMessage().Chat.ID, usersFile)

	msg1, err1 := bot.Send(stats)
	msg2, err2 := bot.Send(users)
	msg3, err3 := bot.Send(phrases)

	MessageDeleter.AddMessage(update_handlers2.NewMessage(&msg1, bot))
	MessageDeleter.AddMessage(update_handlers2.NewMessage(&msg2, bot))
	MessageDeleter.AddMessage(update_handlers2.NewMessage(&msg3, bot))

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

func AddPhraseCallback(params update_handlers2.RedirectedParams) error {
	message := params.Message
	state := params.State

	state.SetState("phrase")
	msg, err := message.EditText(update_handlers2.MessageParams{Text: "отправь фразу", InlineReplyMarkup: &utils.GoBackKeyboard})
	MessageDeleter.AddMessage(msg)
	return err
}
