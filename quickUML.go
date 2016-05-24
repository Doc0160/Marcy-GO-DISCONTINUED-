package main
import (
	"net/url"
	"slack"
)
func qUML(ct *CT, s Slack.OMNI) {
	e, err := explode_cmd(s.Text)
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté!\n"+err.Error())
		return
	}
	if len(e) == 2 {
		Message(ct.Websocket, s, "http://yuml.me/diagram/scruffy/class/"+url.QueryEscape(e[1])+".png")
	} else {
		Message(ct.Websocket, s, "Exemples:\n>http://yuml.me/diagram/scruffy/class/samples")
	}
}
