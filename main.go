package main

import (
	"log"
	"time"

	"github.com/akmyrzza/mechta-notification/internal/config"
	"github.com/akmyrzza/mechta-notification/internal/notifier"
	"github.com/akmyrzza/mechta-notification/internal/repository/sqlite"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// config
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Printf("error to init config: %v", err)
	}

	// repo
	sql, err := sqlite.New(cfg.DB.Name)
	if err != nil {
		log.Printf("error to open databse: %v", err)
	}
	defer sql.Database.Close()

	// bot
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	// notifier
	n := notifier.New(bot, time.Duration(cfg.Notifier.SendInterval), sql)

	// start
	if err := n.Start(); err != nil {
		log.Printf("error to send notification: %v", err)
	}
}
