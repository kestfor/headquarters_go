package notify_service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type User interface {
	UserId() int64
}

type NotifyService struct {
	bot *tgbotapi.BotAPI
}

func NewNotifyService(bot *tgbotapi.BotAPI) *NotifyService {
	return &NotifyService{bot: bot}
}

type NotifyServiceInterface interface {
	Notify(message string, users []User)
}

func (service *NotifyService) Notify(message string, users []User) {
	if service.bot == nil {
		panic("bot arg is nil")
	}

	for _, user := range users {
		messageObj := tgbotapi.NewMessage(user.UserId(), message)
		_, _ = service.bot.Send(messageObj)
	}

}
