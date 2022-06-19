package main

import (
	"time"

	"github.com/go-ping/ping"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checking resource availability via ICMP
func icmpChecker(bot *tgbotapi.BotAPI, config *Config, host string) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}
	pinger.Count = int(config.Icmp.Count)
	pinger.Timeout = 60 * time.Second
	for {
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
		stats := pinger.Statistics()

		if stats.MaxRtt.Seconds() > float64(config.Icmp.Timeout) {
			msg := tgbotapi.NewMessage(config.Telegram.Group, "Host "+host+" ICMP error")
			bot.Send(msg)
		} else if stats.MaxRtt.Milliseconds() > int64(config.Icmp.Timedelay) {
			msg := tgbotapi.NewMessage(config.Telegram.Group, "Host "+host+" ICMP delay is "+stats.MaxRtt.String())
			bot.Send(msg)
		}
		time.Sleep(time.Duration(config.App.Update) * time.Second)
	}
}
