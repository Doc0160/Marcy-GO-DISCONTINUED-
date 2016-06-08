package main
import (
	"strings"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	// "fmt"
)
type(
	RSSChannel struct {
		XMLName xml.Name        `xml:"rss"`
		Items   RSSItems `xml:"channel"`
	}
	RSSItems struct {
		XMLName  xml.Name         `xml:"channel"`
		ItemList []RSSItem `xml:"item"`
	}
	RSSItem struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
	}
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
func implode(c []string, delimiter string)string{
	var a string
	if len(c)==0{
		return ""
	}
	if len(c)==1{
		return c[0]
	}
	for _,v := range c{
		a+=v+delimiter
	}
	if a[:len(a)-1]==delimiter{
		a=a[:len(a)-1]
	}
	return a
}
func explode_cmd(cmd string) ([]string) {
	var r []string
	if len(cmd) > 0 {
		if cmd[0] == '$' || cmd[0] == '%' || cmd[0] == '!' {
			cmd = cmd[1:]
		}
		r = strings.Split(cmd, " ")
		return r
	}
	return r
}
func get_cmd(cmd string)string {
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
			return r
		}else{
			return r
		}
	}
	return r
}
func cut_cmd(cmd string)string {
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
			return r
		}else{
			return r
		}
	}
	return r
}

func commonHTTPRequest(ct *CT, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err!= nil{
		return nil,err
	}else{
		req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
		r, err := ct.Slack.Client.Do(req)
		defer r.Body.Close()
		if err != nil {
			return nil, err
		} else {
			buf, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			} else {
				return buf, nil
			}
		}
	}
}
func commonJsonRequest(ct *CT, url string, o interface{})error{
	println(url)
	buf, err := commonHTTPRequest(ct,url)
	println("apres rq")
	if err != nil {
		return err
	} else {
		println("av dec")
		// println(string(buf))
		err := json.Unmarshal(buf, &o)
		println("apres dec")
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
func commonXMLRequest(ct *CT, url string, o interface{})error{
	println(url)
	buf, err := commonHTTPRequest(ct,url)
	println("apres rq")
	if err != nil {
		return err
	} else {
		println("av dec")
		err := xml.Unmarshal(buf, &o)
		println("apres dec")
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
func reverseString(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)

	return output
}
func flipText(input string) string {
	var flipped string
	var flips = make(map[string]string)

	flips["a"] = "ɐ"
	flips["b"] = "q"
	flips["c"] = "ɔ"
	flips["d"] = "p"
	flips["e"] = "ǝ"
	flips["f"] = "ɟ"
	flips["g"] = "ƃ"
	flips["h"] = "ɥ"
	flips["i"] = "ᴉ"
	flips["j"] = "ɾ"
	flips["k"] = "ʞ"
	flips["l"] = "l"
	flips["m"] = "ɯ"
	flips["n"] = "u"
	flips["o"] = "o"
	flips["p"] = "d"
	flips["q"] = "b"
	flips["r"] = "ɹ"
	flips["s"] = "s"
	flips["t"] = "ʇ"
	flips["u"] = "n"
	flips["v"] = "ʌ"
	flips["w"] = "ʍ"
	flips["x"] = "x"
	flips["y"] = "ʎ"
	flips["z"] = "z"

	flips["A"] = "∀"
	flips["B"] = "B"
	flips["C"] = "Ɔ"
	flips["D"] = "D"
	flips["E"] = "Ǝ"
	flips["F"] = "Ⅎ"
	flips["G"] = "פ"
	flips["H"] = "H"
	flips["I"] = "I"
	flips["J"] = "ſ"
	flips["K"] = "K"
	flips["L"] = "˥"
	flips["M"] = "W"
	flips["N"] = "N"
	flips["O"] = "O"
	flips["P"] = "Ԁ"
	flips["Q"] = "Q"
	flips["R"] = "R"
	flips["S"] = "S"
	flips["T"] = "┴"
	flips["U"] = "∩"
	flips["V"] = "Λ"
	flips["W"] = "M"
	flips["X"] = "X"
	flips["Y"] = "⅄"
	flips["Z"] = "z"

	flips["0"] = "0"
	flips["1"] = "Ɩ"
	flips["2"] = "ᄅ"
	flips["3"] = "Ɛ"
	flips["4"] = "ㄣ"
	flips["5"] = "ϛ"
	flips["6"] = "9"
	flips["7"] = "ㄥ"
	flips["8"] = "8"
	flips["9"] = "6"

	flips[","] = "'"
	flips["."] = "˙"
	flips["?"] = "¿"
	flips["!"] = "¡"
	flips["\""] = ",,"
	flips["'"] = ","
	flips["`"] = ","
	flips["("] = ")"
	flips[")"] = "("
	flips["["] = "]"
	flips["]"] = "["
	flips["{"] = "}"
	flips["}"] = "{"
	flips["<"] = ">"
	flips[">"] = "<"
	flips["&"] = "⅋"
	flips["_"] = "‾"
	
	flips["┻"] = "┬"
	flips["┬"] = "┻"
	flips["━"] = "─"
	flips["─"] = "━"
	
	for _, rune := range input {
		letter := string(rune)
		// get matches
		if flips[letter] != "" {
			flipped += flips[letter]
		} else {
			flipped += letter
		}
	}

	return flipped
}


