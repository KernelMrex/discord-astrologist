package main

import (
	"AstrologyDiscordBot/parser_mail"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func getAstroHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	signsRunes := map[string]rune{
		"aries":       '\u2648',
		"taurus":      '\u2649',
		"gemini":      '\u264A',
		"cancer":      '\u264B',
		"leo":         '\u264C',
		"virgo":       '\u264D',
		"libra":       '\u264E',
		"scorpio":     '\u264F',
		"sagittarius": '\u2650',
		"capricorn":   '\u2651',
		"aquarius":    '\u2652',
		"pisces":      '\u2653',
	}
	signsEmoji := map[string]string{
		"aries":       ":aries:",
		"taurus":      ":taurus:",
		"gemini":      ":gemini:",
		"cancer":      ":cancer:",
		"leo":         ":leo:",
		"virgo":       ":virgo:",
		"libra":       ":libra:",
		"scorpio":     ":scorpius:",
		"sagittarius": ":sagittarius:",
		"capricorn":   ":capricorn:",
		"aquarius":    ":aquarius:",
		"pisces":      ":pisces:",
	}

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "+astro ") {
		return
	}

	sign := ""
	for key, signEmoji := range signsRunes {
		if strings.ContainsRune(m.Content, signEmoji) {
			sign = key
			break
		}
	}
	if sign == "" {
		if _, err := s.ChannelMessageSend(m.ChannelID, "Такого знака зодиака нет"); err != nil {
			Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
			return
		}
		return
	}

	prediction, err := parser_mail.GetAstroForSign(sign)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}

	message := fmt.Sprintf("%v %s\n", signsEmoji[sign], prediction)
	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}
}

func getAllAstroHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	SignsEmoji := map[string]string{
		"aries":       ":aries:",
		"taurus":      ":taurus:",
		"gemini":      ":gemini:",
		"cancer":      ":cancer:",
		"leo":         ":leo:",
		"virgo":       ":virgo:",
		"libra":       ":libra:",
		"scorpio":     ":scorpius:",
		"sagittarius": ":sagittarius:",
		"capricorn":   ":capricorn:",
		"aquarius":    ":aquarius:",
		"pisces":      ":pisces:",
	}

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.TrimSpace(m.Content) != "+astro" {
		return
	}

	predictions, err := parser_mail.GetAstro()
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}

	for sign, prediction := range predictions {
		message := fmt.Sprintf("%s %s\n", SignsEmoji[sign], prediction)
		if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
			Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
			return
		}
	}
}
