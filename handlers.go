package main

import (
	"AstrologyDiscordBot/parser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func getAstroHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	query := strings.TrimSpace(m.Content)
	if !strings.HasPrefix(query, "?astro ") {
		return
	}

	args := strings.Split(query, " ")
	signWord, ok := Env.Signs.GetByPseudo(args[len(args) - 1])
	if !ok {
		if _, err := s.ChannelMessageSend(m.ChannelID, "no such sign"); err != nil {
			Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		}
		return
	}

	prediction, err := parser.GetAstroForSign(signWord)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}

	emoji, ok := Env.Signs.GetEmoji(signWord)
	if !ok {
		if _, err := s.ChannelMessageSend(m.ChannelID, "internal error"); err != nil {
			Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
			return
		}
		return
	}

	message := fmt.Sprintf("%v %s\n", emoji, prediction)
	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}
}

//func getAllAstroHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
//	SignsEmoji := map[string]string{
//		"aries":       ":aries:",
//		"taurus":      ":taurus:",
//		"gemini":      ":gemini:",
//		"cancer":      ":cancer:",
//		"leo":         ":leo:",
//		"virgo":       ":virgo:",
//		"libra":       ":libra:",
//		"scorpio":     ":scorpius:",
//		"sagittarius": ":sagittarius:",
//		"capricorn":   ":capricorn:",
//		"aquarius":    ":aquarius:",
//		"pisces":      ":pisces:",
//	}
//
//	// Ignore all messages created by the bot itself
//	if m.Author.ID == s.State.User.ID {
//		return
//	}
//
//	if strings.TrimSpace(m.Content) != "+astro" {
//		return
//	}
//
//	predictions, err := parser.GetAstro()
//	if err != nil {
//		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
//		return
//	}
//
//	for sign, prediction := range predictions {
//		message := fmt.Sprintf("%s %s\n", SignsEmoji[sign], prediction)
//		if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
//			Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
//			return
//		}
//	}
//}
