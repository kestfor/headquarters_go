package main

import (
	"fmt"
	"headquarters/handlers"
	"headquarters/utils"
	"log"
)
import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var TOKEN = "6696943443:AAEEidx578sqTT3tC0zGg3xppHPeBdVFTFY"

func ChallengeCallback(message handlers.Message, bot *tgbotapi.BotAPI) error {
	_ = message.Delete()
	return message.Answer(handlers.Params{Text: "test", ReplyMarkup: utils.ChallengeReplyKeyboard})
}

func ChooseHouseCallback(message handlers.Message, bot *tgbotapi.BotAPI) error {
	return message.EditText(handlers.Params{Text: "выбери штаб", InlineReplyMarkup: &utils.HousesKeyboard})
}

func GoBackCallback(message handlers.Message, bot *tgbotapi.BotAPI) error {
	return message.EditText(handlers.Params{Text: "главное меню", InlineReplyMarkup: &utils.MenuKeyboard})
}

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}
	callbackManager := handlers.NewCallbackManager(bot)
	callbackManager.RegisterCallback(ChallengeCallback, utils.SEND_LOCATION_INIT)
	callbackManager.RegisterCallback(ChooseHouseCallback, utils.CHOOSE_HOME)
	callbackManager.RegisterCallback(GoBackCallback, utils.MENU)

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.CallbackQuery != nil {
			err = callbackManager.HandleCallback(update)

			// And finally, send a message containing the data received.
			//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			//msg.ReplyMarkup = utils.MenuKeyboard
			//if _, err := bot.Send(msg); err != nil {
			//	panic(err)
			//}
		}

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = utils.MenuKeyboard
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
