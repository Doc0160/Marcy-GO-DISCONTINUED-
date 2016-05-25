package main
import(
	"slack"
	"fmt"
	"golang.org/x/net/websocket" // TODO(doc): use smthg better or custom
)
var marcy Marcy
type Marcy struct{
	cmds Commands
}
//"xoxb-20711630562-YRp4UgH60905Ad4FdKWauigm"
func (m *Marcy)Loop(){
	for true {
		var recv Slack.OMNI
		websocket.JSON.Receive(m.cmds.CT.Websocket, &recv)
		switch recv.Type {
		case "message":
			if _, v := m.cmds.CT.Slack.GetNameById(recv.User); v != "marcy" && len(recv.Text)>0 && recv.Text[0]=='$'{
				e, err := explode_cmd(recv.Text)
				if err==nil{
					fmt.Println(recv)
					if m.cmds.Commands[e[0]] != nil {
						if len(e) > 1 && (e[1] == "h" || e[1] == "help") {
							if m.cmds.Help[e[0]] == "" {
								Message(m.cmds.CT.Websocket, recv, "// TODO: help_text")
							} else {
								Message(m.cmds.CT.Websocket, recv, m.cmds.Help[e[0]])
							}
						} else {
							go m.cmds.Commands[e[0]](&m.cmds.CT, recv)
						}
					} else if e[0] == "h" || e[0] == "help" {
						go func() {
							var t string
							for k, v := range m.cmds.QHelp {
								if v != "" {
									t += "`$" + k + "` : " + v + "\n"
								}
							}
							Message(m.cmds.CT.Websocket, recv, t)
						}()
					} else {
						go default1(&m.cmds.CT, recv)
					}
				}
			}
		case "file_shared":
			fmt.Println(recv.File)
		case "hello":
			println("hello")
		case "presence_change":
			_,v := m.cmds.CT.Slack.GetNameById(recv.User)
			fmt.Println(v, *recv.Presence)
		case "user_typing":
			fmt.Println(recv.User, recv.Channel)
		case "reconnect_url":
			m.cmds.CT.Slack.RTM.URL = recv.URL
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
func (m*Marcy)Handler(n string, f func(*CT, Slack.OMNI), QHelp string, Help string){
	m.cmds.Handler(n, f, QHelp, Help);
}
func (m *Marcy)Init(token string){
	var err error
	m.cmds, err = NewCommands(token)
	if err!=nil{
		println(err.Error())
	}else{
		//fmt.Println(m.cmds.CT.Slack.RTM.URL)
		//
		//
		m.cmds.Handler("exit", func(*CT, Slack.OMNI){
			// TODO(doc): gracefull exit
		},"","")
		//m.cmds.Handler("g", giphy, "giphy", "")
		if err != nil {
			panic(err.Error())
		}
	}
}