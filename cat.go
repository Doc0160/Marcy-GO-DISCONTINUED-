package main
import(
	"./slack"
)
var theCatApiKey = "ODk4MDc"
func cat(ct *CT, s Slack.OMNI){
	Typing(ct.Websocket, s)
	var a cat_api_struct
	err := commonXMLRequest(ct, "http://thecatapi.com/api/images/get?format=xml&results_per_page=1&api_key="+theCatApiKey, &a)
	Typing(ct.Websocket, s)
	if err!=nil{
		Message(ct.Websocket, s, err.Error())
	}else{
		Message(ct.Websocket, s, a.Images[0].URL)
	}
}
type cat_api_struct struct{
	Images []struct{
		URL string `xml:"url"`
		ID string `xml:"id"`
		Source_url string `xml:"source_url"`
	} `xml:"data>images>image"`
}