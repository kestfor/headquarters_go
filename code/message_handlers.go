package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/code/dice"
	"headquarters/code/file_data_base"
	geo2 "headquarters/code/geo"
	"headquarters/code/update_handlers"
	conf "headquarters/code/user_manager"
	utils2 "headquarters/code/utils"
	"log"
	"time"
)

const CHALLENGE_TRY_TIMEOUT time.Duration = 4000000000

func StartCommand(params update_handlers.RedirectedParams) error {
	message := params.Message
	params.State.Clear()

	MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
	MessageDeleter.AddMessage(message)

	err := DataBase.AddUser(&conf.User{UserId: message.ApiMessage.Chat.ID, UserName: message.ApiMessage.From.UserName})
	if err != nil {
		log.Println(err.Error())
	}

	msg, err := message.Answer(update_handlers.MessageParams{Text: "главное меню", ReplyMarkup: &utils2.MenuKeyboard})
	MessageDeleter.AddMessage(update_handlers.NewMessage(&msg, params.Bot))
	return err
}

func StandardMessage(params update_handlers.RedirectedParams) error {
	state := params.State
	message := params.Message

	MessageDeleter.AddMessage(message)
	switch state.GetState() {
	case "":
		return nil
	case "location":
		return locationHandler(message, state, params.Bot)
	case "challenge":
		return challengeHandler(message, state)
	case "phrase":
		return addPhraseHandler(message, state)
	default:
		return nil
	}
}

func challengeHandler(message *update_handlers.Message, state *update_handlers.State) error {

	if message.ApiMessage.Dice == nil {
		return nil
	}

	data := state.GetData().(map[string]any)
	attempts := data["attempts"].(int)
	attempts += 1
	data["attempts"] = attempts

	time.Sleep(CHALLENGE_TRY_TIMEOUT)

	if dice.Collection.Success(*message.ApiMessage.Dice) {

		state.SetState("")
		address := geo2.MainHome

		err := DataBase.AddRecord(&file_data_base.Record{
			UserId:   message.ApiMessage.Chat.ID,
			Time:     time.Now(),
			Address:  address.ToString(),
			Attempts: attempts})

		if err != nil {
			log.Println(err.Error())
		}

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		_, err = message.Answer(update_handlers.MessageParams{Text: "запись добавлена" + utils2.GetRandomHappyEmoji(), ReplyMarkup: &utils2.GoBackKeyboard})
		return err

	} else {

		_, err := message.Answer(update_handlers.MessageParams{Text: utils2.GetRandomChallengePhrase(DataBase.Phrases), ReplyMarkup: &utils2.GoBackKeyboard, ParseMode: "HTML"})
		return err
	}
}

func locationHandler(message *update_handlers.Message, state *update_handlers.State, bot *tgbotapi.BotAPI) error {
	if message.ApiMessage.Location == nil {
		return nil
	}

	latitude := message.ApiMessage.Location.Latitude
	longitude := message.ApiMessage.Location.Longitude

	MessageDeleter.AddMessage(message)

	if geo2.MainHome.Equivalent(geo2.AddressFromLocation(latitude, longitude)) {
		state.SetState("challenge")

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		msg, err := message.Answer(update_handlers.MessageParams{Text: utils2.ChallengeInscription, ReplyMarkup: &utils2.ChallengeReplyKeyboard, ParseMode: "HTML"})
		MessageDeleter.AddMessage(update_handlers.NewMessage(&msg, bot))
		return err
	} else {
		state.SetState("")

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		msg, err := message.Answer(update_handlers.MessageParams{Text: "ты находишься не в том месте" + utils2.GetRandomChallengeEmoji(), ReplyMarkup: &utils2.GoBackKeyboard})
		MessageDeleter.AddMessage(update_handlers.NewMessage(&msg, bot))
		return err
	}

}

func addPhraseHandler(message *update_handlers.Message, state *update_handlers.State) error {
	if message.ApiMessage == nil {
		return nil
	}

	MessageDeleter.AddMessage(message)
	newPhrase := message.ApiMessage.Text

	if newPhrase == "" {
		return nil
	}

	err := DataBase.AddPhrase(newPhrase)

	state.SetState("")

	if err != nil {

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		_, err := message.Answer(update_handlers.MessageParams{Text: "что-то пошло не так", ReplyMarkup: &utils2.GoBackKeyboard})

		return err
	} else {

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		_, err := message.Answer(update_handlers.MessageParams{Text: "фраза добавлена" + utils2.GetRandomHappyEmoji(), ReplyMarkup: &utils2.GoBackKeyboard})
		return err
	}

}
