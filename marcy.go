package main
import(
	"fmt"
	"time"
	"./slack"
	"./TinyJsonDB"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
)
type Marcy struct{
	Commands       Commands
	CT             CT
	DefaultCommand func(*CT, Slack.OMNI)
	HelpCommand    func(*CT, Slack.OMNI,*Commands)
	EventHandlers  map[string]func(*Marcy,Slack.OMNI)
	Master         string
	Me             string
}
type CT struct {
	Websocket  *websocket.Conn
	Slack      Slack.Slack
	TinyJsonDB *TinyJsonDB.TinyJsonDB
	Random     *rand.Rand
	HTTP       *http.Client
}
func NewMarcy(token string, master string, me string)Marcy{
	var m Marcy
	m.Master=master
	m.Me=me
	m.CT.HTTP=&http.Client{
		Timeout: time.Second * 5,
		Jar: nil,
	}
	var err error
	m.CT.TinyJsonDB = TinyJsonDB.New()
	m.CT.Slack.Token = token
	ret, err := m.CT.Slack.API_CALL("rtm.start", nil)
	if err!=nil{
		println("RTM GET URL FAILED")
		panic(err.Error())
	}
	m.CT.Websocket, _, err = websocket.DefaultDialer.Dial(m.CT.Slack.RTM.URL, nil)
	m.CT.Random = rand.New(rand.NewSource(time.Now().Unix()))
	if err!=nil{
		fmt.Println(ret)
		panic(err.Error())
	}else{
		m.Handler("version", func(ct*CT, s Slack.OMNI){
			Message(ct.Websocket, s, "version<1")
		},"version","")
		m.Handler("exit", func(ct*CT, s Slack.OMNI){
			if m.IsMaster(s){
				ct.Websocket.Close()
				// TODO(doc): gracefull exit
				panic("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			}else{
				Message(ct.Websocket, s, "Humain, tu es stupide.")
			}
		},"","")
		m.Alias("quit", "exit");
		m.Alias("ragequit", "exit");
		//
		m.Handler("bonjour", func(ct *CT, s Slack.OMNI){
			Message(ct.Websocket, s, "'Jour.")
		},"","")
		m.Alias("hello","bonjour");
		m.Alias("salut","bonjour");
		m.Handler("aurevoir", func(ct *CT, s Slack.OMNI){
			Message(ct.Websocket, s, "'Rvoir.")
		},"","")
		m.Alias("bye","aurevoir");
		m.Alias("aplus","aurevoir");
		//
		m.Handler("mdr", func(ct *CT, s Slack.OMNI){
			Message(ct.Websocket, s, "HHAHAHAHAHA !")
		},"","")
		m.Alias("lol","mdr");
		
		m.Handler("marcy", func(ct *CT, s Slack.OMNI){
			a := cut_cmd(s.Text)
			Message(ct.Websocket, s, a)
		},"","")
		
		if err != nil {
			panic(err.Error())
		}
		m.HelpCommand=func(CT*CT,recv Slack.OMNI,Commands*Commands) {
			var t string
			for k, v := range *Commands{
				if v.QHelp != "" && v.Alias==false{
					t += "`$" + k + "` : " + v.QHelp + "\n"
				}
			}
			Message(CT.Websocket, recv, t)
		}
	}
	m.EventHandlers = make(map[string]func(*Marcy,Slack.OMNI))
	m.EventHandlers["hello"]=func(*Marcy,Slack.OMNI){
		println("hello")
	}
	m.EventHandlers["reconnect_url"]=func(m *Marcy,recv Slack.OMNI){
		m.CT.Slack.RTM.URL = *recv.URL
	}
	m.EventHandlers["message"]=MessageHandler
	m.EventHandlers["presence_change"]=PresenceChangeHandler
	m.EventHandlers["user_typing"]=DummyHandler
	m.EventHandlers[""]=func(m *Marcy,recv Slack.OMNI){
		println("jkjbkhbk")
		if recv.OK != nil{
			fmt.Println("OK:", *recv.OK, recv.Error)
		}else{
			panic("Got disconnedted")
		}
		println("jkjbkhbk")
	}
	return m;
}
func DummyHandler(m*Marcy, recv Slack.OMNI){}
func MessageHandler(m*Marcy, recv Slack.OMNI){
	if _, v := m.CT.Slack.GetNameById(recv.User); v != "marcy" && len(recv.Text)>0 && recv.Text[0]=='$'{
		fmt.Println(recv.Text)
		e := explode_cmd(recv.Text)
		if m.Commands[e[0]].Command != nil{
			if len(e) > 1 && (e[1] == "h" || e[1] == "help"){
				if m.Commands[e[0]].Help == ""{
					Message(m.CT.Websocket, recv, m.Commands[e[0]].QHelp)
				} else {
					Message(m.CT.Websocket, recv, m.Commands[e[0]].Help)
				}
			} else {
				m.Commands[e[0]].Command(&m.CT, recv)
			}
		} else if e[0] == "h" || e[0] == "help"{
			m.HelpCommand(&m.CT, recv, &m.Commands);
		} else {
			Message(m.CT.Websocket, recv, "Ye ne té comprends pas dins ce qué tû dis !!!\nPour avoir la liste des commandes:\n>`$h`")
		}
		fmt.Println(recv.Text)
	}
}
func PresenceChangeHandler(m*Marcy, recv Slack.OMNI){
	m.CT.Slack.SetPresence(recv.User,*recv.Presence)
	_,v := m.CT.Slack.GetNameById(recv.User)
	if v == "satan_777" && *recv.Presence=="away"{
		var a Slack.OMNI
		a.Channel="G0R8C5KU7"
		Message(m.CT.Websocket, a, ">ICI REPOSE REX\n```Chien aimant```\n```Puceau de premiére```\n```Bot Stupide```\n```Faisait le beau comme personne```\n```Se Léchait les couilles comme personne```")
	}
}
func (m *Marcy)Loop(){
	var recv Slack.OMNI
	for true {
		recv = Slack.OMNI{}
		println("i want to break free")
		websocket.ReadJSON(m.CT.Websocket, &recv)
		println("i want to break free")
		if _, ok := m.EventHandlers[recv.Type]; ok {
			println(recv.Type)
			m.EventHandlers[recv.Type](m, recv)
			println(recv.Type)
		}else{
			fmt.Println("What :",recv)
		}
	}
}
func (m*Marcy)Handler(n string, f func(*CT, Slack.OMNI), QHelp string, Help string){
	if m.Commands == nil {
		m.Commands = make(map[string]Command)
	}
	m.Commands[n] = Command{
		Command : f,
		QHelp:    QHelp,
		Help:     Help,
		Alias:    false,
	}
}
func(m*Marcy)Alias(n2 string, n string){
	m.Commands[n2] = Command{
		Command : m.Commands[n].Command,
		QHelp:    m.Commands[n].QHelp,
		Help:     m.Commands[n].Help,
		Alias:    true,
	}
}
func(m*Marcy)AliasMulti(n2 []string, n string){
	for _,v := range n2{
		m.Commands[v] = Command{
			Command : m.Commands[n].Command,
			QHelp:    m.Commands[n].QHelp,
			Help:     m.Commands[n].Help,
			Alias:    true,
		}
	}
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
	websocket.WriteJSON(ws, Slack.Typing{
		ID:      time.Now().String(),
		Type:    "typing",
		Channel: s.Channel,
	})
}
// Send a message event in the channel specfied in the incomming message
func Message(ws *websocket.Conn, s Slack.OMNI, text string) {
	websocket.WriteJSON(ws, Slack.Message{
		ID:      time.Now().String(),
		Type:    "message",
		Channel: s.Channel,
		Text:    text,
	})
}