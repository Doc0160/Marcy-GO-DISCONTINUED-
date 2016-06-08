package main
import (
	"./slack"
	"time"
	// "fmt"
)
func timestamp(ct *CT, s Slack.OMNI){
	Message(ct.Websocket, s, time.Now().String())
}