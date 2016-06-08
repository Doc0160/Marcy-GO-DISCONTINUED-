package main
import(
	"./slack"
)
type Commands map[string]Command
type Command struct{
	Command func(*CT, Slack.OMNI)
	QHelp   string
	Help    string
	Alias   bool
	Hidden  bool
}
func(c*Command)GetHelp()string{
	return ""
}
func(c*Command)GetQHelp()string{
	return ""
}