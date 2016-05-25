package main
import (
	"math/rand"
	"github.com/Doc0160/Marcy/slack"
	"strings"
	"time"
	"golang.org/x/net/html"
)
func codinglove(ct *CT, s Slack.OMNI) {
	rand.Seed(time.Now().Unix())
	var (
		title      string
		img        string
		has_title  bool = false
		has_img    bool = false
		grab_title bool = false
		// err   error
	)
	r, err := commonHttpRequest(ct, "http://lesjoiesducode.fr/random")
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	} else {
		func() {
			z := html.NewTokenizer(*r)
			for {
				tt := z.Next()
				switch {
				case tt == html.ErrorToken:
					// End of the document, we're done
					if has_img && has_title {
						Message(ct.Websocket, s, title+"\n"+img)
						return
					}
				case tt == html.StartTagToken:
					t := z.Token()
					switch t.Data {
					case "h1":
						for _, a := range t.Attr {
							if a.Key == "class" && a.Val == "blog-post-title" {
								grab_title = true
							}
						}
					case "img":
						for _, a := range t.Attr {
							if a.Key == "src" {
								img = a.Val
								has_img = true
							}
						}
					}
				case tt == html.TextToken:
					if grab_title && !has_title {
						has_title = true
						t := z.Token()
						title = strings.Trim(t.Data, " \t\r\n")
					}
				}
			}
			(*r).Close()
		}()
	}
}
