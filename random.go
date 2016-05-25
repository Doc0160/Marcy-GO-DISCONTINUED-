package main
import (
	"github.com/Doc0160/Marcy/slack"
	"time"
	// "fmt"
)
func timestamp(ct *CT, s Slack.OMNI){
	Message(ct.Websocket, s, time.Now().String())
}