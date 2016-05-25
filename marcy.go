package main
import(
	"fmt"
	"time"
	"github.com/Doc0160/Marcy/slack"
	"github.com/Doc0160/Marcy/TinyJsonDB"
	"golang.org/x/net/websocket" // TODO(doc): use smthg better or custom
	"math/rand"
)
type Marcy struct{
	//cmds Commands
	Commands map[string]Command
	CT CT
}
type Command struct{
	Command func(*CT, Slack.OMNI)
	QHelp   string
	Help    string
}
type CT struct {
	Websocket  *websocket.Conn
	Slack      Slack.Slack
	TinyJsonDB *TinyJsonDB.TinyJsonDB
	Random     *rand.Rand
}
func NewMarcy(token string)Marcy{
	var m Marcy
	var err error
	m.CT.TinyJsonDB = TinyJsonDB.New()
	m.CT.Slack.Token = token
	_, err = m.CT.Slack.API_CALL("rtm.start", nil)
	if err!=nil{
		panic(err.Error())
	}
	m.CT.Websocket, err = websocket.Dial(m.CT.Slack.RTM.URL, "", "https://slack.com/")
	m.CT.Random = rand.New(rand.NewSource(time.Now().Unix()))
	if err!=nil{
		panic(err.Error())
	}else{
		m.Handler("exit", func(*CT, Slack.OMNI){
			// TODO(doc): gracefull exit
		},"","")
		if err != nil {
			panic(err.Error())
		}
	}
	return m;
}
func (m *Marcy)Loop(){
	for true {
		var recv Slack.OMNI
		websocket.JSON.Receive(m.CT.Websocket, &recv)
		switch recv.Type{
		case "message":
			if _, v := m.CT.Slack.GetNameById(recv.User); v != "marcy" && len(recv.Text)>0 && recv.Text[0]=='$'{
				e, err := explode_cmd(recv.Text)
				if err==nil{
					fmt.Println(recv)
					if m.Commands[e[0]].Command != nil{
						if len(e) > 1 && (e[1] == "h" || e[1] == "help"){
							if m.Commands[e[0]].Help == ""{
								Message(m.CT.Websocket, recv, m.Commands[e[0]].QHelp)
							} else {
								Message(m.CT.Websocket, recv, m.Commands[e[0]].Help)
							}
						} else {
							go m.Commands[e[0]].Command(&m.CT, recv)
						}
					} else if e[0] == "h" || e[0] == "help"{
						go func() {
							var t string
							for k, v := range m.Commands{
								if v.QHelp != ""{
									t += "`$" + k + "` : " + v.QHelp + "\n"
								}
							}
							Message(m.CT.Websocket, recv, t)
						}()
					} else {
						go default1(&m.CT, recv)
					}
				}
			}
		case "file_shared":
			fmt.Println(recv.File)
		case "hello":
			println("hello")
		case "presence_change":
			_,v := m.CT.Slack.GetNameById(recv.User)
			fmt.Println(v, *recv.Presence)
		case "user_typing":
			fmt.Println(recv.User, recv.Channel)
		case "reconnect_url":
			m.CT.Slack.RTM.URL = recv.URL
			// fmt.Println(recv.URL)
		case "":
			if recv.OK != nil && *recv.OK == false {
				fmt.Println("OK:", recv.OK, recv.Error)
			}
		default:
			fmt.Println(recv)
		}
	}
}
func (m*Marcy)Handler(n string, f func(*CT, Slack.OMNI), QHelp string, Help string){
	if m.Commands == nil {
		m.Commands = make(map[string]Command)
	}
	m.Commands[n] = Command{
		Command : f,
		QHelp: QHelp,
		Help: Help,
	}
}
func(m*Marcy)Alias(n2 string, n string){
	m.Commands[n2]=m.Commands[n]
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