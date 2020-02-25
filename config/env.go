package config

import (
	"AstrologyDiscordBot/zodiac"
	"log"
)

type Environment struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Config      *Configuration
	Signs       *zodiac.Associate
}
