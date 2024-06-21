package utils

import "strings"

const (
	MENU               = "menu"
	SEND_LOCATION_INIT = "send_location_init"
	CHOOSE_HOME        = "choose_home"
	ADD_PHRASE_INIT    = "add_phrase_init"
	DOWNLOAD_STAT      = "download_stat"
)

func GetCallbackData(rawString string) string {
	return strings.Replace(rawString, " ", "_", -1)
}
