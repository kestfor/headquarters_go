package main

import (
	"headquarters/dice"
	"headquarters/file_data_base"
	"headquarters/geo"
	"headquarters/handlers"
	conf "headquarters/user_manager"
	"headquarters/utils"
	"log"
	"time"
)

const CHALLENGE_TRY_TIMEOUT time.Duration = 4000000000

func StartCommand(params handlers.Params) error {
	message := params.Message

	err := DataBase.AddUser(&conf.User{UserId: message.ApiMessage.Chat.ID, UserName: message.ApiMessage.From.UserName})
	if err != nil {
		log.Println(err.Error())
	}

	return message.Answer(handlers.FunctionParams{Text: "главное меню", ReplyMarkup: &utils.MenuKeyboard})
}

func StandardMessage(params handlers.Params) error {
	state := params.State
	message := params.Message

	switch state.GetState() {
	case "":
		return nil
	case "location":
		return locationHandler(message, state)
	case "challenge":
		return challengeHandler(message, state)
	default:
		return nil
	}
}

func challengeHandler(message *handlers.Message, state *handlers.State) error {

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

		return message.Answer(handlers.FunctionParams{Text: "запись добавлена", ReplyMarkup: &utils.GoBackKeyboard})

	} else {

		return message.Answer(handlers.FunctionParams{Text: "анлак, пробуй еще", ReplyMarkup: &utils.GoBackKeyboard})

	}
}

func locationHandler(message *handlers.Message, state *handlers.State) error {
	if message.ApiMessage.Location == nil {
		return nil
	}

	latitude := message.ApiMessage.Location.Latitude
	longitude := message.ApiMessage.Location.Longitude

	if geo.MainHome.Equivalent(geo.AddressFromLocation(latitude, longitude)) {
		state.SetState("challenge")
		_ = message.Delete()

		return message.Answer(handlers.FunctionParams{Text: utils.ChallengeInscription, ReplyMarkup: &utils.ChallengeReplyKeyboard, ParseMode: "HTML"})

	} else {
		state.SetState("")

		_ = message.Delete()

		return message.Answer(handlers.FunctionParams{Text: "ты находишься не в том месте", ReplyMarkup: &utils.GoBackKeyboard})
	}

}
