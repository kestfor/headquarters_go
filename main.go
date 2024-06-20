package main

import (
	"headquarters/file_data_base"
	"headquarters/geo"
	"headquarters/handlers"
	"headquarters/utils"
	"log"
)

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var TOKEN = "6696943443:AAEEidx578sqTT3tC0zGg3xppHPeBdVFTFY"
var DataBase *file_data_base.DataBase

func main() {
	var err error
	DataBase, err = file_data_base.NewDataBase("data/users.json", "data/stats.json")

	if err != nil {
		log.Panic(err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
		return
	}

	handler := handlers.NewHandler(bot)

	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfDima)
	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfIlya)
	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfAlena)

	handler.CallbackManager.RegisterCallback(DownloadFiles, utils.DOWNLOAD_STAT)
	handler.CallbackManager.RegisterCallback(ChallengeCallback, utils.SEND_LOCATION_INIT)
	handler.CallbackManager.RegisterCallback(ChooseHouseCallback, utils.CHOOSE_HOME)
	handler.CallbackManager.RegisterCallback(GoBackCallback, utils.MENU)
	handler.CommandManager.RegisterCommand(StartCommand, "start")
	handler.CommandManager.RegisterCommand(StartCommand, "menu")
	handler.CommandManager.RegisterCommand(StandardMessage, "")

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go func() {
			err := handler.HandleUpdate(update)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}
