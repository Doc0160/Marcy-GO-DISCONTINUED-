package main
import(
	"github.com/Doc0160/Marcy/slack"
)
type Command struct{
	Command func(*CT, Slack.OMNI)
	QHelp   string
	Help    string
	Alias   bool
}
func(c*Command)GetHelp()string{
	return ""
}
func(c*Command)GetQHelp()string{
	return ""
}