package main
import(
	"./slack"
)
var loremS []string = []string{
	"loreo",
	"ip",
	"linux",
	"geek",
	"nerd",
	"machine",
	"learning",
	"chat",
	"bot",
	"php",
	"goroutine",
	"golang",
	"c++",
	"pizza",
	"coca",
	"c",
	"fuck",
	"câble",
	"ananas",
	"github",
}
func loreo(ct *CT, s Slack.OMNI){
	var a string
	var nb int = ct.Random.Intn(10)
	for i:=0;i<nb;i++{
		a+=loremS[ct.Random.Intn(len(loremS))]+" "
	}
	Message(ct.Websocket, s, a)
}