package main
import (
	"errors"
	"io"
	"net/http"
	"strings"
	"encoding/json"
)
func toX(a string, b int, p string, s bool) string {
	for len(a) < b {
		if s {
			a = p + a
		} else {
			a += p
		}
	}
	return a
}
func implode(c []string, delimiter byte)string{
	var a string
	if len(c)==0{
		return ""
	}
	if len(c)==1{
		return c[0]
	}
	for _,v := range c{
		a+=v+string(delimiter)
	}
	if a[:len(a)-1]==string(delimiter){
		a=a[:len(a)-1]
	}
	return a
}
func explode_cmd(cmd string) ([]string, error) {
	var r []string
	if len(cmd) > 0 {
		if cmd[0] == '$' || cmd[0] == '%' || cmd[0] == '!' {
			cmd = cmd[1:]
		}
			r = strings.Split(cmd, " ")
			return r, nil
		/*}else{
			return r, errors.New("Not a command")
		}*/
	}
	return r, errors.New("Texte vide")
}
func get_cmd(cmd string) (string, error) {
	var r string
	if len(cmd) > 0 {
		if cmd[0] == '$' || cmd[0] == '%' || cmd[0] == '!' {
			cmd = cmd[1:]
			i:=strings.IndexAny(cmd, " ")
			if i==-1{
				r=cmd
			}else{
				r = cmd[:i]
			}
			return r, nil
		}else{
			return r, errors.New("Not a command")
		}
	}
	return r, errors.New("Texte vide")
}
func cut_cmd(cmd string) (string, error) {
	var r string
	if len(cmd) > 0 {
		if cmd[0] == '$' || cmd[0] == '%' || cmd[0] == '!' {
			cmd = cmd[1:]
			i:=strings.IndexAny(cmd, " ")
			if i==-1{
				r=""
			}else{
				r = cmd[i+1:]
			}
			return r, nil
		}else{
			return r, errors.New("Not a command")
		}
	}
	return r, errors.New("Texte vide")
}
func commonHttpRequest(ct *CT, url string) (*io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
	r, err := ct.Slack.Client.Do(req)
	if err != nil {
		return nil, err
		//		Message(ct.Websocket, s, "Y'a une couille dans le paté !\n"+err.Error())
	} else {
		return &r.Body, err
		//		r.Body.Close()
	}
}
func commonJsonRequest(ct *CT, url string, o interface{})error{
	r, err := commonHttpRequest(ct, url)
	if err != nil {
		return err
	} else {
		err := json.NewDecoder(*r).Decode(&o)
		defer (*r).Close()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}