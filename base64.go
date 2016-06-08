package main
import(
	"./slack"
	b64 "encoding/base64"
)
func base64(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	Message(ct.Websocket, s, b64.StdEncoding.EncodeToString([]byte(s.Text)))
}
func unbase64(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	r, err := b64.StdEncoding.DecodeString(s.Text)
	if err!=nil{
		Message(ct.Websocket, s, err.Error())
	}else{
		Message(ct.Websocket, s, string(r))
	}
}