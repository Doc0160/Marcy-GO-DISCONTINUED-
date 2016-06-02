package main
import(
	"github.com/Doc0160/Marcy/slack"
	// "fmt"
)
func hn(ct *CT, s Slack.OMNI){
	var a RSSChannel
	var f []Slack.Attachment
	Typing(ct.Websocket,s)
	commonXMLRequest(ct, "https://news.ycombinator.com/rss", &a)
	i:=0
	for _,v := range a.Items.ItemList{
		if i>15{
			break
		}
		f=append(f, Slack.Attachment{
			Title: v.Title,
			TitleLink: v.Link,
			Fallback: "HackerNews: Bunch of News",
		})
		i++
	}
	Typing(ct.Websocket,s)
	_, err := ct.Slack.API_CALL("chat.postMessage", map[string]interface{}{
		"as_user":     "true",
		"channel":     s.Channel,
		"attachments": f,
	})
	if err != nil {
		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	}
}