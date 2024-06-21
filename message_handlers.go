package main

import (
	"headquarters/dice"
	"headquarters/file_data_base"
	"headquarters/geo"
	updHnd "headquarters/update_handlers"
	conf "headquarters/user_manager"
	"headquarters/utils"
	"log"
	"time"
)

const CHALLENGE_TRY_TIMEOUT time.Duration = 4000000000

func StartCommand(params updHnd.RedirectedParams) error {
	message := params.Message
	params.State.Clear()

	err := DataBase.AddUser(&conf.User{UserId: message.ApiMessage.Chat.ID, UserName: message.ApiMessage.From.UserName})
	if err != nil {
		log.Println(err.Error())
	}

	return message.Answer(updHnd.MessageParams{Text: "главное меню", ReplyMarkup: &utils.MenuKeyboard})
}

func StandardMessage(params updHnd.RedirectedParams) error {
	state := params.State
	message := params.Message

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

func challengeHandler(message *updHnd.Message, state *updHnd.State) error {

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
		address := geo.MainHome

		err := DataBase.AddRecord(&file_data_base.Record{
			UserId:   message.ApiMessage.Chat.ID,
			Time:     time.Now(),
			Address:  address.ToString(),
			Attempts: attempts})

		if err != nil {
			log.Println(err.Error())
		}

		return message.Answer(updHnd.MessageParams{Text: "запись добавлена" + utils.GetRandomHappyEmoji(), ReplyMarkup: &utils.GoBackKeyboard})

	} else {

		return message.Answer(updHnd.MessageParams{Text: utils.GetRandomChallengePhrase(DataBase.Phrases), ReplyMarkup: &utils.GoBackKeyboard, ParseMode: "HTML"})

	}
}

func locationHandler(message *updHnd.Message, state *updHnd.State) error {
	if message.ApiMessage.Location == nil {
		return nil
	}

	latitude := message.ApiMessage.Location.Latitude
	longitude := message.ApiMessage.Location.Longitude

	if geo.MainHome.Equivalent(geo.AddressFromLocation(latitude, longitude)) {
		state.SetState("challenge")
		_ = message.Delete()

		return message.Answer(updHnd.MessageParams{Text: utils.ChallengeInscription, ReplyMarkup: &utils.ChallengeReplyKeyboard, ParseMode: "HTML"})

	} else {
		state.SetState("")

		_ = message.Delete()

		return message.Answer(updHnd.MessageParams{Text: "ты находишься не в том месте" + utils.GetRandomChallengeEmoji(), ReplyMarkup: &utils.GoBackKeyboard})
	}

}

func addPhraseHandler(message *updHnd.Message, state *updHnd.State) error {
	if message.ApiMessage == nil {
		return nil
	}

	newPhrase := message.ApiMessage.Text

	if newPhrase == "" {
		return nil
	}

	err := DataBase.AddPhrase(newPhrase)

	state.SetState("")

	if err != nil {
		return message.Answer(updHnd.MessageParams{Text: "что-то пошло не так", ReplyMarkup: &utils.GoBackKeyboard})
	} else {
		return message.Answer(updHnd.MessageParams{Text: "фраза добавлена" + utils.GetRandomHappyEmoji(), ReplyMarkup: &utils.GoBackKeyboard})
	}

}
