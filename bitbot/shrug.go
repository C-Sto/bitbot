package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var ShrugTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!shrug"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `¯\_(ツ)_/¯`)
		return true
	},
}
