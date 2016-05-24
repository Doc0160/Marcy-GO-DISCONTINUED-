package main
import (
	"math/rand"
	"slack"
	"time"
)
var magic8 = []string{
	"« Essaye plus tard »",
	"« Essaye encore »",
	"« Pas d'avis »",
	"« C'est ton destin »",
	"« Le sort en est jeté »",
	"« Une chance sur deux »",
	"« Repose ta question »",

	"« D'après moi oui »",
	"« C'est certain »",
	"« Oui absolument »",
	"« Tu peux compter dessus »",
	"« Sans aucun doute »",
	"« Très probable »",
	"« Oui »",
	"« C'est bien parti »",

	"« C'est non »",
	"« Peu probable »",
	"« Faut pas rêver »",
	"« N'y compte pas »",
	"« Impossible »",
}
func doMagic8ball(ct *CT, s Slack.OMNI){
	rand.Seed(time.Now().Unix())
	q,_ := cut_cmd(s.Text)
	var text = magic8[rand.Intn(len(magic8))]
	var att Slack.Attachment
	att.Text = text
	att.Fallback = text
	att.Title = "Magic 8 Ball"
	att.Color = "#000000"
	att.ThumbURL = "https://33.media.tumblr.com/avatar_ed2e9fed4447_128.png"
	Typing(ct.Websocket, s)
	_, err := ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
		"as_user": "true",
		"channel": s.Channel,
		"text":    q+" ",
		"attachments": []Slack.Attachment{
			att,
		},
	})
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	}
}