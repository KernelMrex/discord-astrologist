package main

import (
	"AstrologyDiscordBot/config"
	"AstrologyDiscordBot/zodiac"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Env *config.Environment

func init() {
	// Loggers
	errorLogger := log.New(os.Stderr, "ERROR | ", log.Lshortfile|log.Ltime)
	infoLogger := log.New(os.Stdout, "INFO  | ", log.Lshortfile|log.Ltime)

	// Config
	configPath := os.Getenv("BOT_ASTROLOGIST_CONFIG")
	cfg, err := config.LoadConfigFromJsonFile(configPath)
	if err != nil {
		errorLogger.Fatalln("[ init ]", err)
	}

	// Zodiac
	signsMap := os.Getenv("BOT_ASTROLOGIST_SIGNS")
	signs, err := zodiac.LoadAssociateFromJson(signsMap)
	if err != nil {
		errorLogger.Fatalln("[ init ]", err)
	}

	// Env
	Env = &config.Environment{
		ErrorLogger: errorLogger,
		InfoLogger:  infoLogger,
		Config:      cfg,
		Signs:       signs,
	}
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Env.Config.BotConfig.Secret)
	if err != nil {
		Env.ErrorLogger.Fatalln("[ main ]", "error creating Discord session,", err)
	}

	// Register the playUserSoundHandler func as a callback for MessageCreate events.
	dg.AddHandler(getAstroHandler)

	// Open a websocket connection to Discord and begin listening.
	if err := dg.Open(); err != nil {
		Env.ErrorLogger.Fatalln("[ main ]", "error opening connection,", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	Env.InfoLogger.Println("[ main ]", "Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	if err := dg.Close(); err != nil {
		Env.ErrorLogger.Fatalln("[ main ]", "error closing:", err)
	}
}
