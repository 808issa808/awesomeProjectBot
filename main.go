package main

import (
	"awesomeProjectBot/telegram"
	"log"
	"time"

	_ "github.com/lib/pq"
	tele "gopkg.in/tucnak/telebot.v2"
)

func main() {
	pref, err := tele.NewBot(tele.Settings{
		Token:  "5408087287:AAG9IBfgHEo6Fe-ZDZx_xFcI8jwxvWL8OXg",
		Poller: &tele.LongPoller{Timeout: 15 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b := telegram.NewBot(pref)
	b.Init()

	b.Bot.Start()
}
