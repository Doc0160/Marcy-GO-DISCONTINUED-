package main
import(
	// "math/rand"
	// "time"
	"./slack"
	// "fmt"
	// "encoding/json"
)
func giphy(ct *CT, s Slack.OMNI){
	var gs giphy_struct
	s.Text = cut_cmd(s.Text)
	temp := explode_cmd(s.Text)
	if len(temp)==0{
		Message(ct.Websocket, s, "Usage:\n>$g `something`")
		return
	}
	s.Text = implode(temp, "+")
	err := commonJsonRequest(ct, "https://api.giphy.com/v1/gifs/search?q="+s.Text+"&api_key=dc6zaTOxFJmzC", &gs)
	if err != nil {
		Message(ct.Websocket, s, "Paté: "+err.Error())
	} else {
		if len(gs.Data)>0{
			Message(ct.Websocket, s, gs.Data[ct.Random.Intn(len(gs.Data))].Images.Original.URL)
		}else{
			Message(ct.Websocket, s, "-none found-")
		}
	}
}

type giphy_struct struct {
	Data []struct {
		// Type string `json:"type"`
		// ID string `json:"id"`
		// Slug string `json:"slug"`
		// URL string `json:"url"`
		// BitlyGifURL string `json:"bitly_gif_url"`
		// BitlyURL string `json:"bitly_url"`
		// EmbedURL string `json:"embed_url"`
		// Username string `json:"username"`
		// Source string `json:"source"`
		// Rating string `json:"rating"`
		// ContentURL string `json:"content_url"`
		// SourceTld string `json:"source_tld"`
		// SourcePostURL string `json:"source_post_url"`
		// ImportDatetime string `json:"import_datetime"`
		// TrendingDatetime string `json:"trending_datetime"`
		Images struct {
			Original struct {
				URL string `json:"url"`
				Width string `json:"width"`
				Height string `json:"height"`
				Size string `json:"size"`
				Frames string `json:"frames"`
				Mp4 string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Webp string `json:"webp"`
				WebpSize string `json:"webp_size"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
	/*Meta struct {
		Status int `json:"status"`
		Msg string `json:"msg"`
	} `json:"meta"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count int `json:"count"`
		Offset int `json:"offset"`
	} `json:"pagination"`*/
}