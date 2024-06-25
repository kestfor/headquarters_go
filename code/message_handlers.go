package main

import (
	"headquarters/code/dice"
	"headquarters/code/file_data_base"
	geo2 "headquarters/code/geo"
	"headquarters/code/notify_service"
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

	err := DataBase.AddUser(conf.NewTelegramUser(message.ApiMessage.Chat.ID, message.ApiMessage.From.UserName))
	if err != nil {
		log.Println(err.Error())
	}

	msg, err := message.Answer(update_handlers.MessageParams{Text: "главное меню", ReplyMarkup: &utils2.MenuKeyboard})
	MessageDeleter.AddMessage(msg)
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
		return locationHandler(message, state)
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
		address := geo2.MainHome.Address
		loc, _ := time.LoadLocation("Asia/Krasnoyarsk")
		record := file_data_base.Record{
			UserId:   message.ApiMessage.Chat.ID,
			Time:     time.Now().In(loc),
			Address:  address.ToString(),
			Attempts: attempts}

		err := DataBase.AddRecord(&record)

		if err != nil {
			log.Println(err.Error())
		}

		recordOwner := DataBase.GetUser(record.UserId)
		allUsers := DataBase.Users()

		var usersToNotify = make([]notify_service.User, 0, len(allUsers)-1)
		for _, user := range allUsers {
			if user.UserId() != recordOwner.UserId() {
				usersToNotify = append(usersToNotify, &user)
			}
		}

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		_, err = message.Answer(update_handlers.MessageParams{Text: "запись добавлена" + utils2.GetRandomHappyEmoji(), ReplyMarkup: &utils2.GoBackKeyboard})
		go NotifyService.Notify("@"+recordOwner.UserName()+" зачекинился у @"+geo2.MainHome.Owner, usersToNotify)
		return err

	} else {

		_, err := message.Answer(update_handlers.MessageParams{Text: utils2.GetRandomChallengePhrase(DataBase.Phrases), ReplyMarkup: &utils2.GoBackKeyboard, ParseMode: "HTML"})
		return err
	}
}

func locationHandler(message *update_handlers.Message, state *update_handlers.State) error {
	if message.ApiMessage.Location == nil {
		return nil
	}

	latitude := message.ApiMessage.Location.Latitude
	longitude := message.ApiMessage.Location.Longitude

	MessageDeleter.AddMessage(message)

	if geo2.MainHome.Address.Equivalent(geo2.AddressFromLocation(latitude, longitude)) {
		state.SetState("challenge")

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		msg, err := message.Answer(update_handlers.MessageParams{Text: utils2.ChallengeInscription, ReplyMarkup: &utils2.ChallengeReplyKeyboard, ParseMode: "HTML"})
		MessageDeleter.AddMessage(msg)
		return err
	} else {
		state.SetState("")

		MessageDeleter.DeleteMessages(message.ApiMessage.Chat.ID)
		msg, err := message.Answer(update_handlers.MessageParams{Text: "ты находишься не в том месте" + utils2.GetRandomChallengeEmoji(), ReplyMarkup: &utils2.GoBackKeyboard})
		MessageDeleter.AddMessage(msg)
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
