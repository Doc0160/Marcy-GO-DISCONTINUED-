package main
import (
	// "math/rand"
	"fmt"
	"slack"
	"time"
	"golang.org/x/net/websocket"
)
func main() {
	var err error
	cmds,err := NewCommands("xoxb-20711630562-PRAus6HxcY8Cxl42xrjf3INi")
	if err!=nil{
		
	}else{
		fmt.Println(cmds.CT.Slack.RTM.URL)
		//
		var xkcd = newXkcd(cmds.CT.TinyJsonDB, &cmds.CT)
		//
		cmds.Handler("exit", func(*CT, Slack.OMNI){
			// TODO(doc): gracefull exit
		},"","")
		cmds.Handler("g", giphy, "giphy", "")
		cmds.Handler("ping", ping, "pong!", "PONG!")
		cmds.Handler("debug", doDebug, "debug", "debug")
		cmds.Handler("xkcd", xkcd.do_xkcd, "XKCD !", "")
		cmds.Handler("timestamp", timestamp, "", "")
		cmds.Handler("pokedex", do_pkdx, "Pokemon !", "")
		cmds.Alias("pokedex", "pkdx")
		cmds.Handler("perv", perv, "Pervers", "")
		cmds.Handler("fch", do_forecastHourly, "La prévision météo des heures qui viennent", "`$fch reims`\n`$fch paris`") // TOUPDATE
		cmds.Handler("meme", memify, "//TODO DESC", "")
		cmds.Handler("m8b", doMagic8ball, "Magic 8 ball, pose ta question et la boule te répondera.", "")
		cmds.Handler("jdc", codinglove, "Les joies du code", "")
		cmds.Handler("wf", doWarframeAlert, "Warframe Alerts", "")
		cmds.Handler("treta", treta, "Treta", "")
		cmds.Handler("cat", cat, "Chats <3", "")
		cmds.Handler("quml", qUML, "quick UML: http://yuml.me/diagram/scruffy/class/samples", "")
		cmds.Handler("prosit", func(ct *CT, s Slack.OMNI) {
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
			Message(cmds.CT.Websocket, s, classe[ct.Random.Intn(len(classe))])
		}, "Totalement pas copié de \"(tars/case) nouveau prosit\".", "")
		if err != nil {
			panic(err.Error())
		}
		for true {
			var recv Slack.OMNI
			websocket.JSON.Receive(cmds.CT.Websocket, &recv)
			switch recv.Type {
			case "message":
				if _, v := cmds.CT.Slack.GetNameById(recv.User); v != "marcy" && recv.Text[0]=='$'{
					e, err := explode_cmd(recv.Text)
					if err == nil {
						fmt.Println(recv)
						if cmds.Commands[e[0]] != nil {
							if len(e) > 1 && (e[1] == "h" || e[1] == "help") {
								if cmds.Help[e[0]] == "" {
									Message(cmds.CT.Websocket, recv, "// TODO: help_text")
								} else {
									Message(cmds.CT.Websocket, recv, cmds.Help[e[0]])
								}
							} else {
								go cmds.Commands[e[0]](&cmds.CT, recv)
							}
						} else if e[0] == "h" || e[0] == "help" {
							go func() {
								var t string
								for k, v := range cmds.QHelp {
									if v != "" {
										t += "`$" + k + "` : " + v + "\n"
									}
								}
								Message(cmds.CT.Websocket, recv, t)
							}()
						} else {
							go default1(&cmds.CT, recv)
						}
					}
				}
			case "file_shared":
				fmt.Println(recv.File)
			case "hello":
				println("hello")
			case "presence_change":
				_,v := cmds.CT.Slack.GetNameById(recv.User)
				fmt.Println(v, *recv.Presence)
			case "user_typing":
				fmt.Println(recv.User, recv.Channel)
			case "reconnect_url":
				cmds.CT.Slack.RTM.URL = recv.URL
				// fmt.Println(recv.URL)
			case "":
				if recv.OK != nil && *recv.OK == false {
					fmt.Println("OK:", recv.OK, recv.Error)
				} /*else {
					println("All good !")
				}*/
			default:
				fmt.Println(recv)
			}
		}
	}
}
// Send a typing event in the channel specfied in the incomming message
func Typing(ws *websocket.Conn, s Slack.OMNI) {
	websocket.JSON.Send(ws, Slack.Typing{
		ID:      time.Now().String(),
		Type:    "typing",
		Channel: s.Channel,
	})
}
// Send a message event in the channel specfied in the incomming message
func Message(ws *websocket.Conn, s Slack.OMNI, text string) {
	websocket.JSON.Send(ws, Slack.Message{
		ID:      time.Now().String(),
		Type:    "message",
		Channel: s.Channel,
		Text:    text,
	})
}