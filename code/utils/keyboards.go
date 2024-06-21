package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/code/dice"
	"headquarters/code/geo"
)

var ChallengeReplyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(dice.FOOTBALL),
		tgbotapi.NewKeyboardButton(dice.BASKETBALL),
		tgbotapi.NewKeyboardButton(dice.BOWLING),
	),
)

var GeolocationReplyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("отправить геолокацию"),
	),
)

var MenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("зачекиниться", SEND_LOCATION_INIT)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("выбрать штаб", CHOOSE_HOME)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("добавить фразу", ADD_PHRASE_INIT)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("выгрузить статистику", DOWNLOAD_STAT)),
)

var GoBackKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("назад", MENU)),
)

//func getHousesButtons() []tgbotapi.InlineKeyboardButton {
//	var buttons []tgbotapi.InlineKeyboardButton
//	for key := range geo.Houses {
//		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(key, key))
//	}
//	return buttons
//}

var HousesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(geo.HomeOfAlena, geo.HomeOfAlena)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(geo.HomeOfIlya, geo.HomeOfIlya)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(geo.HomeOfDima, geo.HomeOfDima)),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("назад", MENU)),
)
