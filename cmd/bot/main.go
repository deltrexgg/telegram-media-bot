/*
	Package main is the entry point for the Telegram media bot application.
	It loads configuration, initializes the bot, and manages application lifecycle.
	main loads configuration, initializes the Telegram bot
*/

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	config "github.com/deltrexgg/telegram-media-bot/internal/config"
	"github.com/deltrexgg/telegram-media-bot/internal/handler"
	"github.com/deltrexgg/telegram-media-bot/internal/repository"
	"github.com/deltrexgg/telegram-media-bot/internal/service"
	"github.com/deltrexgg/telegram-media-bot/internal/telegram"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file")
	}

	config.InitDB()

	apiKey := os.Getenv("TELEGRAM_API")
	if apiKey == "" {
		log.Fatal("TELEGRAM_API is not set")
	}

	folderrepo := repository.FolderRepoMethod()
	folderservice := service.FolderServiceMethod(folderrepo)
	folderhandler := handler.FolderHandler(folderservice)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	opts := []bot.Option{
		bot.WithDefaultHandler(folderhandler.CreateFolder),
	}

	b, err := telegram.InitBot(apiKey, opts...)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Telegram bot started")
	b.Start(ctx)
	log.Println("Telegram bot stopped")
}
