package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/code/file_data_base"
	"headquarters/code/geo"
	"headquarters/code/notify_service"
	"headquarters/code/update_handlers"
	"headquarters/code/utils"
	"log"
	"os"
)

var TEST_TOKEN = "6696943443:AAEEidx578sqTT3tC0zGg3xppHPeBdVFTFY"
var TOKEN = "6729922230:AAGeyBOYUVERijSeabmwjs352XMNJWtICQ8"
var DataBase *file_data_base.DataBase
var MessageDeleter *Deleter

var NotifyService *notify_service.NotifyService

func main() {
	var err error
	_ = os.Mkdir("data", os.ModePerm)
	DataBase, err = file_data_base.NewDataBase("data/users.json", "data/stats.json", "data/phrases.txt")
	MessageDeleter = NewDeleter()

	if err != nil {
		log.Panic(err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(TEST_TOKEN)
	NotifyService = notify_service.NewNotifyService(bot)

	if err != nil {
		log.Panic(err)
		return
	}

	handler := update_handlers.NewHandler(bot)

	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfDima)
	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfIlya)
	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfAlena)
	handler.CallbackManager.RegisterCallback(HousesCallback, geo.HomeOfAnton)

	handler.CallbackManager.RegisterCallback(AddPhraseCallback, utils.ADD_PHRASE_INIT)
	handler.CallbackManager.RegisterCallback(DownloadFiles, utils.DOWNLOAD_STAT)
	handler.CallbackManager.RegisterCallback(ChallengeCallback, utils.SEND_LOCATION_INIT)
	handler.CallbackManager.RegisterCallback(ChooseHouseCallback, utils.CHOOSE_HOME)
	handler.CallbackManager.RegisterCallback(GoBackCallback, utils.MENU)
	handler.CommandManager.RegisterCommand(StartCommand, "start")
	handler.CommandManager.RegisterCommand(StartCommand, "menu")
	handler.CommandManager.RegisterCommand(StandardMessage, "")

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
