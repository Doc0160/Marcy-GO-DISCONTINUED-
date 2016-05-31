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
	Commands       map[string]Command
	CT             CT
	DefaultCommand func(*CT, Slack.OMNI)
	HelpCommand    func(*CT, Slack.OMNI,*map[string]Command)
	//
	Master string
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
func NewMarcy(token string, master string)Marcy{
	var m Marcy
	m.Master=master
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
		m.Handler("exit", func(ct*CT, s Slack.OMNI){
			if m.IsMaster(s){
				// TODO(doc): gracefull exit
				panic("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}else{
				Message(ct.Websocket, s, "Humain, tu es stupide.")
			}
		},"","")
		m.Alias("quit", "exit");
		
		m.Handler("mdr", func(ct *CT, s Slack.OMNI){
			Message(ct.Websocket, s, "HHAHAHAHAHA !")
		},"","")
		
		m.Alias("lol","mdr");
		if err != nil {
			panic(err.Error())
		}
		m.HelpCommand=func(CT*CT,recv Slack.OMNI,Commands*map[string]Command) {
			var t string
			for k, v := range *Commands{
				if v.QHelp != ""{
					t += "`$" + k + "` : " + v.QHelp + "\n"
				}
			}
			Message(CT.Websocket, recv, t)
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
						go m.HelpCommand(&m.CT, recv, &m.Commands);
					} else {
						go default1(&m.CT, recv)
					}
				}
			}
		case "file_shared":
			fmt.Println(*recv.File)
			Message(m.CT.Websocket,Slack.OMNI{Channel:"D0LM5HH25"},recv.File.URLPrivateDownload)
			//D0LM5HH25
		case "hello":
			println("hello")
		case "presence_change":
			m.CT.Slack.SetPresence(recv.User,*recv.Presence)
			_,v := m.CT.Slack.GetNameById(recv.User)
			fmt.Println(v, *recv.Presence)
			if v == "satan_test777" && *recv.Presence=="away"{
				var a Slack.OMNI
				a.Channel="G0R8C5KU7"
				Message(m.CT.Websocket, a, ">ICI REPOSE REX\n```Chien aimant```\n```Puceau de premiére```\n```Bot Stupide```\n```Faisait le beau comme personne```\n```Se Léchait les couilles comme personne```")
			}
		case "user_typing":
			fmt.Println(recv.User, recv.Channel)
		case "reconnect_url":
			m.CT.Slack.RTM.URL = *recv.URL
		case "":
			if recv.OK != nil{
				fmt.Println("OK:", *recv.OK, recv.Error)
			}else{
				fmt.Println(recv)
				panic("euuuuuuuuuuuuuuuuh")
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
func(m*Marcy)SetDefaultCommand(f func(*CT,Slack.OMNI)){
	m.DefaultCommand=f
}
func(m*Marcy)IsMaster(recv Slack.OMNI)bool{
	_, v := m.CT.Slack.GetNameById(recv.User)
	return  v == m.Master
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