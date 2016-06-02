package main
import (
	"net/url"
	"github.com/Doc0160/Marcy/slack"
)
func qUML(ct *CT, s Slack.OMNI) {
	e := explode_cmd(s.Text)
	if len(e) == 2 {
		Message(ct.Websocket, s, "http://yuml.me/diagram/scruffy/class/"+url.QueryEscape(e[1])+".png")
	} else {
		Message(ct.Websocket, s, "Exemples:\n>http://yuml.me/diagram/scruffy/class/samples")
	}
}
