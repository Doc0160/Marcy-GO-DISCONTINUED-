package main
import (
	"math/rand"
	"./slack"
	"strings"
	"time"
	"bytes"
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
		grab_title2 bool = false
		// err   error
	)
	buf, err := commonHTTPRequest(ct,"http://thecodinglove.com/random")
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	} else {
		func() {
			z := html.NewTokenizer(bytes.NewReader(buf))
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
					case "div":
						for _, a := range t.Attr {
							if a.Key == "class" && a.Val == "centre" {
								grab_title = true
							}
						}
					case "h3":
						if grab_title {
							grab_title2 = true
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
					if grab_title2 && !has_title {
						has_title = true
						t := z.Token()
						title = strings.Trim(t.Data, " \t\r\n")
					}
				}
			}
		}()
	}
}
func joieducode(ct *CT, s Slack.OMNI) {
	rand.Seed(time.Now().Unix())
	var (
		title      string
		img        string
		has_title  bool = false
		has_img    bool = false
		grab_title bool = false
		// err   error
	)
	buf, err := commonHTTPRequest(ct,"http://lesjoiesducode.fr/random")
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	} else {
		func() {
			z := html.NewTokenizer(bytes.NewReader(buf))
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
		}()
	}
}
