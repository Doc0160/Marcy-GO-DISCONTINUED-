package main
import(
	"github.com/Doc0160/Marcy/slack"
	"encoding/xml"
	"fmt"
	// "strconv"
)
var theCatApiKey = "ODk4MDc"
var CatCache StackString
func cat(ct *CT, s Slack.OMNI){
	if CatCache.Size()>0{
		a, _ := CatCache.Pop()
		Message(ct.Websocket, s, a)
	}else{
		Typing(ct.Websocket, s)
		r, err := commonHttpRequest(ct, "http://thecatapi.com/api/images/get?format=xml&results_per_page=1&api_key="+theCatApiKey)
		if err!=nil{
			Message(ct.Websocket, s, err.Error())
		}else{
			Typing(ct.Websocket, s)
			var a cat_api_struct
			err := xml.NewDecoder(*r).Decode(&a)
			defer (*r).Close()
			if err!=nil{
				Message(ct.Websocket, s, err.Error())
			}else{
				Message(ct.Websocket, s, a.Images[0].URL)
			}
		}
	}
	if CatCache.Size()==0{
		go CatPreload(ct, 5)
	}
}
func getACat(ct *CT)(string,error){
	r, err := commonHttpRequest(ct, "http://thecatapi.com/api/images/get?format=xml&results_per_page=1&api_key="+theCatApiKey)
	fmt.Println(*r)
	if err!=nil{
		return "", err
	}else{
		var a cat_api_struct
		err := xml.NewDecoder(*r).Decode(&a)
		defer (*r).Close()
		if err!=nil{
			return "", err
		}else{
			return a.Images[0].URL,nil
		}
	}
}
func CatPreload(ct *CT, x int){
	var str string
	var err error
	for CatCache.Size()<x{
		str, err = getACat(ct)
		if err==nil{
			CatCache.Push(str)
		}
	}
}
type cat_api_struct struct{
	Images []struct{
		URL string `xml:"url"`
		ID string `xml:"id"`
		Source_url string `xml:"source_url"`
	} `xml:"data>images>image"`
}