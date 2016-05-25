package main

import (
	"fmt"
	"runtime/debug"
	"github.com/Doc0160/Marcy/slack"
	// "golang.org/x/net/websocket"
)

func default1(ct *CT, s Slack.OMNI) {
	Message(ct.Websocket, s, "Ye ne té comprends pas dins ce qué tû dis !!!\nPour avoir la liste des commandes:\n>`$h`")
}
func ping(ct *CT, s Slack.OMNI) {
	Message(ct.Websocket, s, "Pong !")
}
func doDebug(ct *CT, s Slack.OMNI) {
	var gcs debug.GCStats
	debug.ReadGCStats(&gcs)
	Message(ct.Websocket, s, fmt.Sprintf("%+v\n%+v", gcs, s))
}
