package main
import (
	// "fmt"
	"TinyJsonDB"
	"math/rand"
	"github.com/Doc0160/Marcy/slack"
	"time"
	"golang.org/x/net/websocket"
)
func NewCommands(token string)(Commands,error){
	var err error
	c:= Commands{}
	c.CT.TinyJsonDB = TinyJsonDB.New()
	c.CT.Slack.Token = token
	_, err = c.CT.Slack.API_CALL("rtm.start", nil)
	if err!=nil{
		return c,err
	}
	println(c.CT.Slack.RTM.URL)
	c.CT.Websocket, err = websocket.Dial(c.CT.Slack.RTM.URL, "", "https://slack.com/")
	c.CT.Random = rand.New(rand.NewSource(time.Now().Unix()))
	return c,err
}
type Command struct{
	Command func(*CT, Slack.OMNI)
	QHelp   string
	Help    map[string]string
}
type Commands struct {
	Commands map[string]func(*CT, Slack.OMNI)
	QHelp    map[string]string
	Help     map[string]string
	CT       CT
}
type CT struct {
	Websocket  *websocket.Conn
	Slack      Slack.Slack
	TinyJsonDB *TinyJsonDB.TinyJsonDB
	Random     *rand.Rand
}
func (c *Commands) Handler(n string, f func(*CT, Slack.OMNI), QHelp string, Help string) {
	if c.Commands == nil {
		c.Commands = make(map[string]func(*CT, Slack.OMNI))
		c.QHelp = make(map[string]string)
		c.Help = make(map[string]string)
	}
	c.Commands[n] = f
	c.QHelp[n] = QHelp
	c.Help[n] = Help
}
func (c *Commands) Alias(n2 string, n string) {
	c.Commands[n] = c.Commands[n2]
	c.QHelp[n] = ""
	// c.QHelp[n]=c.QHelp[n2]
	c.Help[n] = c.Help[n2]
}