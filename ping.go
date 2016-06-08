package main

import (
	"fmt"
	"runtime/debug"
	"./slack"
	// "golang.org/x/net/websocket"
)
func ping(ct *CT, s Slack.OMNI) {
	Message(ct.Websocket, s, "Pong !")
}
func doDebug(ct *CT, s Slack.OMNI) {
	var gcs debug.GCStats
	debug.ReadGCStats(&gcs)
	Message(ct.Websocket, s, fmt.Sprintf("%+v\n%+v", gcs, s))
}
