package main
import (
	"github.com/Doc0160/Marcy/slack"
)
func main() {
	marcy := NewMarcy("key", "doc0160")
	//
	var xkcd = newXkcd(marcy.CT.TinyJsonDB, &marcy.CT)
	marcy.Handler("g", giphy, "giphy", "");
	marcy.Handler("flip", flip, "", "");
	marcy.Handler("ping", ping, "pong!", "PONG!")
	marcy.Handler("debug", doDebug, "debug", "debug")
	marcy.Handler("xkcd", xkcd.do_xkcd, "XKCD !", "")
	marcy.Handler("timestamp", timestamp, "", "")
	marcy.Handler("pokedex", do_pkdx, "Pokemon !", "")
	marcy.Alias("pkdx", "pokedex")
	marcy.Handler("perv", perv, "Pervers", "")
	marcy.Alias("girl", "perv")
	marcy.Alias("girls", "perv")
	marcy.Handler("bonlundi", perv_get_boys, "Perverse", "")
	marcy.Alias("needboys", "bonlundi")
	marcy.Alias("boy", "bonlundi")
	marcy.Alias("boys", "bonlundi")
	marcy.Handler("fch", do_forecastHourly, "La prévision météo des heures qui viennent", "`$fch reims`\n`$fch paris`") // TOUPDATE
	marcy.Handler("meme", memify, "//TODO DESC", "")
	marcy.Handler("m8b", doMagic8ball, "Magic 8 ball, pose ta question et la boule te répondera.", "")
	marcy.Alias("mb", "m8b")
	marcy.Alias("m8", "m8b")
	marcy.Alias("8b", "m8b")
	marcy.Alias("magic8ball", "m8b")
	marcy.Alias("magic8", "m8b")
	marcy.Handler("jdc", joieducode, "les Joies Du Code", "")
	marcy.Handler("tcl", codinglove, "The Coding Love", "")
	marcy.Handler("wf", doWarframeAlert, "Warframe Alerts", "")
	marcy.Handler("treta", treta, "Treta", "")
	marcy.Handler("roll", roll, "", "")
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
		Message(ct.Websocket, s, classe[ct.Random.Intn(len(classe))])
	}, "Totalement pas copié de \"(tars/case) nouveau prosit\".", "")
	marcy.Handler("caresse", func(ct *CT, s Slack.OMNI) {
		Message(ct.Websocket, s, "MAIS CH'UIS PAS UN CHIEN MOI !\n_(tu peux gratter un peu plus vers la droite ?)_")
	}, "", "")
	marcy.Handler("rex", func(ct *CT, s Slack.OMNI) {
		Message(ct.Websocket, s, "REX!")
		Message(ct.Websocket, s, "!calme toi")
		Message(ct.Websocket, s, "!carresse")
		Message(ct.Websocket, s, "!fait le beau")
		Message(ct.Websocket, s, "!c'est un gentil rex ça")
		Message(ct.Websocket, s, "!c'est pour qui la babale")
		Message(ct.Websocket, s, "!c'est pour qui la babale")
		Message(ct.Websocket, s, "!va cherhcer")
		Message(ct.Websocket, s, "attends qu'il parte et va se cacher")
	}, "", "")
	marcy.Handler("hn", hn, "Hacker news : https://news.ycombinator.com/","")
	//
	marcy.Loop()
}