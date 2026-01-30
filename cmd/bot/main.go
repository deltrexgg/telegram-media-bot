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
	"github.com/deltrexgg/telegram-media-bot/migrations"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file")
	}

	// return the DB access pointer to a var in the main file which is passed to repo of the modules
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database failed to connect")
	}

	migrations.AutoMigrate(db)

	apiKey := os.Getenv("TELEGRAM_API")
	if apiKey == "" {
		log.Fatal("TELEGRAM_API is not set")
	}

	folderrepo := repository.NewFolderRepo(db)
	folderservice := service.NewFolderService(folderrepo)
	folderhandler := handler.NewFolderHandler(folderservice)

	filerepo := repository.NewFileRepo(db)
	fileservice := service.NewFileService(filerepo)
	filehandler := handler.NewFileHandler(fileservice)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	opts := []bot.Option{
		bot.WithMessageTextHandler("/makefolder", bot.MatchTypePrefix, folderhandler.CreateFolder),
		bot.WithMessageTextHandler("/folders", bot.MatchTypePrefix, folderhandler.GetFolderList),
		bot.WithMessageTextHandler("/addfiles", bot.MatchTypePrefix, folderhandler.GetFolderList),
		bot.WithMessageTextHandler("/stop", bot.MatchTypePrefix, filehandler.StopShare),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, filehandler.CallBackHandler),
		bot.WithDefaultHandler(filehandler.CommonFilehandler),
	}

	b, err := telegram.InitBot(apiKey, opts...)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Telegram bot started")
	b.Start(ctx)
	log.Println("Telegram bot stopped")
}
