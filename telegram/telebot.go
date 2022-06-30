package telegram

import (
	tele "gopkg.in/tucnak/telebot.v2"
)

var Auth = &tele.ReplyMarkup{ResizeReplyKeyboard: true}
var BtnAuthCafe = Auth.Text("Login As Admin of Cafe")
var BtnAuthCompany = Auth.Text("Login As Admin of Company")

var menuAdmin = &tele.ReplyMarkup{ResizeReplyKeyboard: true}
var BtnNewMenu = menuAdmin.Text("Create new menu:")

type TeleBot struct {
	Bot *tele.Bot
}

func NewBot(bot *tele.Bot) *TeleBot {
	return &TeleBot{Bot: bot}
}

func (b *TeleBot) Init() {
	b.initMenus()
	b.Registration()
	b.adminMenu()
}

func (b *TeleBot) initMenus() {
	Auth.Reply(
		Auth.Row(BtnAuthCafe),
		Auth.Row(BtnAuthCompany),
	)
	menuAdmin.Reply(
		Auth.Row(BtnNewMenu),
	)

}
