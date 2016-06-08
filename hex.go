package main
import(
	"./slack"
	h "encoding/hex"
)
func hex(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	Message(ct.Websocket, s, h.EncodeToString([]byte(s.Text)))
}
func unhex(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	r, err := h.DecodeString(s.Text)
	if err!=nil{
		Message(ct.Websocket, s, err.Error())
	}else{
		Message(ct.Websocket, s, string(r))
	}
}