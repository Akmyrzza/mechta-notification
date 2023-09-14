package notifier

import (
	"fmt"
	"time"

	"github.com/akmyrzza/mechta-notification/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Notifier struct {
	bot          *tgbotapi.BotAPI
	sendInterval time.Duration
	repository   repository.Repository
}

func New(bot *tgbotapi.BotAPI, sendInerval time.Duration, repository repository.Repository) *Notifier {
	return &Notifier{
		bot:          bot,
		sendInterval: sendInerval,
		repository:   repository,
	}
}

func (n *Notifier) Start() error {
	ticker := time.NewTicker(n.sendInterval)
	defer ticker.Stop()

	channels, err := n.repository.GetAllChannels()
	if err != nil {
		return err
	}

	workers, err := n.repository.GetAll()
	if err != nil {
		return err
	}

	for _, channel := range channels {
		err := n.SendMsg(channel.TelegramId, workers)
		if err != nil {
			return err
		}
	}

	//здесь не придумано как сделать
	//приложение не закрывается пока не будет ошибки
	//подход новый пока в раздумьях

	// errorCh := make(chan error)

	// for _, channel := range channels {
	// 	go func(TelegramId int, workers []string) {
	// 		err := n.SendMsg(TelegramId, workers)
	// 		if err != nil {
	// 			errorCh <- err
	// 			return
	// 		}
	// 	}(channel.TelegramId, workers)
	// }

	// select {
	// case errCh := <-errorCh:
	// 	log.Printf("Error while sending messages: %v", errCh)
	// }

	return nil
}

func (n *Notifier) SendMsg(channelId int, workers []string) error {

	for _, worker := range workers {
		tgMsg := tgbotapi.NewMessage(int64(channelId), fmt.Sprintf("happy birthday %s!\n", worker))

		_, err := n.bot.Send(tgMsg)
		if err != nil {
			return err
		}
	}

	return nil
}
