package tgbot

import (
	"context"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var (
	botCli   *tgbotapi.BotAPI
	initOnce = sync.Once{}
)

const (
	alertChatId    = int64(-1002143524628)
	techPushChatId = int64(-1002441901462)
	botToken       = "7859518429:AAH_Qnzyyr_dMJ22NHAq-WFPwlo7D6ufwhw"
)

func ensureTG() {

	initOnce.Do(func() {
		bot, err := tgbotapi.NewBotAPI(botToken)
		if err != nil {
			logrus.WithContext(context.Background()).
				Errorf("Init TgBot Faield: %v\n", err)
			return
		}

		bot.Debug = false

		logrus.WithContext(context.Background()).
			Infof("TgBot Authorized on account %s\n", bot.Self.UserName)
		botCli = bot
	})

}

func SendAlert(text string) {
	ensureTG()
	_, _ = botCli.Send(tgbotapi.NewMessage(
		alertChatId, text))
}

func SendTechAlert(text string) {
	ensureTG()
	_, _ = botCli.Send(tgbotapi.NewMessage(
		techPushChatId, text))
}
