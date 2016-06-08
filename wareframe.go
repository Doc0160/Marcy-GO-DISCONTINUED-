package main
import (
	"./slack"
	"github.com/SlyMarbo/rss"
)
func doWarframeAlert(ct *CT, s Slack.OMNI) {
	wf, _ := warframeAlert()
	var t string
	for _, v := range wf.Items {
		t += v.Title + "\n"
	}
	Message(ct.Websocket, s, t)
}
func warframeAlert() (*rss.Feed, error) {
	feed, err := rss.Fetch("http://content.warframe.com/dynamic/rss.php")
	if err != nil {
		return feed, err
	}
	return feed, nil
}
