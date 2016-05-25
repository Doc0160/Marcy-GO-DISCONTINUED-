package main
import (
	// "fmt"
	URL "net/url"
	"github.com/Doc0160/Marcy/slack"
	"strconv"
)
var meme_templates = map[string]string{
	"tentacule":                "http://d37zvtkb7mfi0s.cloudfront.net/wp-content/uploads/2015/04/tentacles-tentacle-rape-cat-girl-hentai-wet-pussy-breasts-tits-blush-molested-hentai.jpg",
	"tentacules":               "http://d37zvtkb7mfi0s.cloudfront.net/wp-content/uploads/2015/04/tentacles-tentacle-rape-cat-girl-hentai-wet-pussy-breasts-tits-blush-molested-hentai.jpg",
	"php":                      "https://upload.wikimedia.org/wikipedia/commons/c/c1/PHP_Logo.png",
	"bonlundi":                 "https://36.media.tumblr.com/feb87d7d2e0ffd05010d9d2ad7604766/tumblr_nmjt3ows5x1s9j7tro1_1280.jpg",
	"bonlundi1":                "https://40.media.tumblr.com/a418e1c33e18b508f6bc9a49bc2e1324/tumblr_nvgpguuxIO1qzjwylo1_1280.png",
	"boobs":                    "http://www.hentai-sex-and-toons.com/thumbs/3693/9.jpg",
	"aliensguy":                "aliensguy.jpg",
	"alienguy":                 "aliensguy.jpg",
	"braceyourself":            "braceyourself.jpg",
	"confessionbear":           "confessionbear.jpg",
	"condescendingwonka":       "condescendingwonka.jpg",
	"firstdayinternetkid":      "firstdayinternetkid.jpg",
	"grumpycat":                "grumpycat.jpg",
	"onedoesnotsimply":         "onedoesnotsimply.jpg",
	"scumbagsteve":             "scumbagsteve.jpg",
	"wankil_bis":               "https://pbs.twimg.com/media/CgwYxXNWYAA2gKY.jpg",
	"wankil":                   "https://lh3.googleusercontent.com/-cDzVNi4OGxY/U3jvQK1UH4I/AAAAAAAAAI8/51uTImdxWJA/s630-fcrop64=1,5d710000f47aff90/1500x500.png",
	"overlyattachedgirlfriend": "overlyattachedgirlfriend.jpg",
	"goodguygreg":              "goodguygreg.jpg",
	"ermahgerd":                "ermahgerd.jpg",
	"conspiracykeanu":          "conspiracykeanu.jpg",
	"yunoguy":                  "yunoguy.jpg",
	"sadfrog":                  "http://memegen.link/sadfrog/_/_.jpg",
	"xzibit":                   "xzibit.jpg",
	"successkid":               "successkid.jpg",
	"futuramafry":              "futuramafry.jpg",
	"firstworldproblems":       "firstworldproblems.jpg",
	"buzz":                     "http://memegen.link/buzz/_/_.jpg",
	"xxeverywhere":             "http://memegen.link/buzz/_/_.jpg",
}
func memify(ct *CT, s Slack.OMNI) {
	e, _ := explode_cmd(s.Text)
	if len(e) == 1 {
		var t string
		for k, _ := range meme_templates {
			t += "`" + k + "`\n"
		}
		Message(ct.Websocket, s, t)
	} else if len(e) < 5 && len(e) > 1 {
		if len(e) < 3 {
			e = append(e, "_")
		}
		if len(e) < 4 {
			e = append(e, "_")
		}
		Message(ct.Websocket, s, "http://memeifier.com/"+URL.QueryEscape(e[2])+"/"+URL.QueryEscape(e[3])+"/"+(URL.QueryEscape(meme_templates[e[1]])))
	} else {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+strconv.Itoa(len(e)))
	}
}