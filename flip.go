package main
import(
	"github.com/Doc0160/Marcy/slack"
)
func flip(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	if s.Text==""{
		s.Text="┬─┬"
	}
	Message(ct.Websocket, s, "(╯°□°）╯︵"+flipText(reverseString(s.Text)))
}
func unflip(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	if s.Text==""{
		s.Text="┻━┻"
	}
	Message(ct.Websocket, s, s.Text+"ノ( º _ ºノ) ")
}