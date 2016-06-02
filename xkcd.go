package main
import (
	// "fmt"
	"github.com/Doc0160/Marcy/TinyJsonDB"
	"encoding/json"
	"net/http"
	"github.com/Doc0160/Marcy/slack"
	"strconv"
)
func newXkcd(tjdb *TinyJsonDB.TinyJsonDB, ct *CT) *XKCD {
	x := XKCD{
		Last: 1677,
		TJDB: tjdb,
	}
	go x.StoreLast(ct)
	return &x
}
func (x *XKCD) StoreLast(ct *CT) {
	var a XKCDcomic
	r, err := commonHttpRequest(ct, "https://xkcd.com/info.0.json")
	err = json.NewDecoder(*r).Decode(&a)
	if !x.TJDB.IsSetTable("xkcd"){
		x.TJDB.CreateTable("xkcd")
	}
	x.TJDB.Data["xkcd"]["last"] = a.Num
	(*r).Close()
	if err != nil {
		// return a, err
		return
	} else {
		// return a, nil
		return
	}
}
func (x *XKCD) do_xkcd(ct *CT, s Slack.OMNI) {
	e := explode_cmd(s.Text)
	if len(e) > 0 {
		var nb int
		var err error
		if len(e) == 1 {
			nb = ct.Random.Intn(ct.TinyJsonDB.Data["xkcd"]["last"].(int))
			err = nil
		} else {
			nb, err = strconv.Atoi(e[1])
		}
		if err == nil {
			if nb == -1 {
				nb = ct.TinyJsonDB.Data["xkcd"]["last"].(int)
				go x.StoreLast(ct)
			}
			Typing(ct.Websocket, s)
			xc, err := x.xkcd(nb)
			Typing(ct.Websocket, s)
			if err == nil {
				ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
					"as_user": "true",
					"channel": s.Channel,
					"attachments": []Slack.Attachment{
						Slack.Attachment{
							Title:     xc.Title + " (" + strconv.Itoa(xc.Num) + ")",
							TitleLink: "https://xkcd.com/" + strconv.Itoa(xc.Num),
							Text:      xc.Alt,
							ImageURL:  xc.Img,
						},
					},
				})
			} else {
				if nb > ct.TinyJsonDB.Data["xkcd"]["last"].(int) || nb < 1 {
					Message(ct.Websocket, s, "Le numéro que vous essayez de joindre n'existe pas ...")
				} else {
					Message(ct.Websocket, s, "Y'a une couille dans le paté !")
				}
			}
		} else {
			Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
		}
	} else {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !")
	}
}
func (x *XKCD) xkcd(num int) (XKCDcomic, error) {
	client := &http.Client{}
	var a XKCDcomic
	req, err := http.NewRequest("GET", "https://xkcd.com/"+strconv.Itoa(num)+"/info.0.json", nil)
	req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
	r, err := client.Do(req)
	if err != nil {
		return a, err
	} else {
		err := json.NewDecoder(r.Body).Decode(&a)
		defer r.Body.Close()
		if err != nil {
			return a, err
		} else {
			return a, nil
		}
	}
}
type XKCDcomic struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}
type XKCD struct {
	Last int
	TJDB *TinyJsonDB.TinyJsonDB
}