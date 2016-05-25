package main
import (
	// "math/rand"
	// "fmt"
	"slack"
	"time"
	"golang.org/x/net/websocket"
)
func main() {
	marcy := NewMarcy("xoxb-20711630562-YRp4UgH60905Ad4FdKWauigm")
	//
	var xkcd = newXkcd(marcy.cmds.CT.TinyJsonDB, &marcy.cmds.CT)
	marcy.Handler("g", giphy, "giphy", "");
	marcy.Handler("ping", ping, "pong!", "PONG!")
	marcy.Handler("debug", doDebug, "debug", "debug")
	marcy.Handler("xkcd", xkcd.do_xkcd, "XKCD !", "")
	marcy.Handler("timestamp", timestamp, "", "")
	marcy.Handler("pokedex", do_pkdx, "Pokemon !", "")
	//marcy.Alias("pokedex", "pkdx")
	marcy.Handler("perv", perv, "Pervers", "")
	marcy.Handler("fch", do_forecastHourly, "La prévision météo des heures qui viennent", "`$fch reims`\n`$fch paris`") // TOUPDATE
	marcy.Handler("meme", memify, "//TODO DESC", "")
	marcy.Handler("m8b", doMagic8ball, "Magic 8 ball, pose ta question et la boule te répondera.", "")
	marcy.Handler("jdc", codinglove, "Les joies du code", "")
	marcy.Handler("wf", doWarframeAlert, "Warframe Alerts", "")
	marcy.Handler("treta", treta, "Treta", "")
	marcy.Handler("cat", cat, "Chats <3", "")
	marcy.Handler("quml", qUML, "quick UML: http://yuml.me/diagram/scruffy/class/samples", "")
	marcy.Handler("prosit", func(ct *CT, s Slack.OMNI) {
		var classe = []string{
			"Lucas",
			"Maxence(tentacules)",
			"Romain",
			"Alice",
			"Maxence(pas tentacules)",
			"Thomas",
			"Joshua",
			"Gaelle",
			"Romain",
		}
		Message(marcy.cmds.CT.Websocket, s, classe[ct.Random.Intn(len(classe))])
	}, "Totalement pas copié de \"(tars/case) nouveau prosit\".", "")
	//
	marcy.Loop()
}