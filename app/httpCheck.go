package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checking the availability of the site via the HTTP protocol
func httpCheck(bot *tgbotapi.BotAPI, config *Config, site struct {
	Url      string
	Elements []string
}) {
	client := http.Client{
		Timeout: time.Duration(config.App.Update) * time.Second,
	}
	for {
		errorHTML := 0
		deface := false
		for i := 0; i < int(config.Http.Repeat); i++ {

			// Check site for available
			resp, err := client.Get(site.Url)
			if err != nil {
				msg := tgbotapi.NewMessage(config.Telegram.Group, "Site "+site.Url+" HTTP get error")
				bot.Send(msg)
			}
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				msg := tgbotapi.NewMessage(config.Telegram.Group, "Site "+site.Url+" HTTP get error")
				bot.Send(msg)
			}
			body := string(bodyBytes)

			// Check for OK status
			if resp.StatusCode != 200 {
				msg := tgbotapi.NewMessage(config.Telegram.Group, "Site "+site.Url+" HTTP error. Code "+strconv.Itoa(resp.StatusCode))
				bot.Send(msg)
				break
			}
			
			// Check element on site (defacing)
			for _, element := range site.Elements {
				if !strings.Contains(body, element) {
					msg := tgbotapi.NewMessage(config.Telegram.Group, "Site "+site.Url+" defaced. Element '"+element+"' not found.")
					bot.Send(msg)
					deface = true
				}
			}
			if deface {
				break
			}
		}
		if errorHTML >= int(config.Http.Repeat-1) {
			msg := tgbotapi.NewMessage(config.Telegram.Group, "Site "+site.Url+" HTTP get error")
			bot.Send(msg)
		}
		time.Sleep(time.Duration(config.Telegram.Group) * time.Second)
	}
}
