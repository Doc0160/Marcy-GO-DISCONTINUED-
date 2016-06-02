package main
import(
	"github.com/Doc0160/Marcy/slack"
	"fmt"
	// "strings"
)
type token struct{
	Name  string
	Value string
	Arg1  *token
	Arg2  *token
}
func roll(ct *CT, s Slack.OMNI){
	s.Text = cut_cmd(s.Text)
	a := explode_cmd(s.Text)
	z,e := roll_tokenizer(a[0])
	Message(ct.Websocket, s, fmt.Sprintf("%+v\n%+v", z,e))
}

func roll_tokenizer(a string)([]string,[]string){
	// var tokens token
	// var stack stackynator
	var b []string
	var t []string
	for i:=0;i<len(a);i++{
		switch a[i]{
			case '0','1','2','3','4','5','6','7','8','9':
				var c string
				for i<len(a) && (a[i]=='0'||a[i]=='1'||a[i]=='2'||a[i]=='3'||a[i]=='4'||a[i]=='5'||a[i]=='6'||a[i]=='7'||a[i]=='8'||a[i]=='9') {
					c+=string(a[i])
					i++
				}
				i--
				b=append(b,c)
				t=append(t,"NUMBER")
			case 'd':
				b=append(b,string(a[i]))
				t=append(t,"DICE")
			case '+':
				b=append(b,string(a[i]))
				t=append(t,"ADD")
			case '-':
				b=append(b,string(a[i]))
				t=append(t,"SUB")
			case '*':
				b=append(b,string(a[i]))
				t=append(t,"MULT")
			case '/':
				b=append(b,string(a[i]))
				t=append(t,"DIV")
			case '%':
				b=append(b,string(a[i]))
				t=append(t,"MOD")
			default:
				b=append(b,string(a[i]))
				t=append(t,"_")
		}
	}
	return b,t
}