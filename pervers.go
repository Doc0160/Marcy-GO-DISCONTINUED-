package main
import (
	"./slack"
)
var perv_girls = []string{
	"YogaPants",
	"boobs",
	"SexyButNotPorn",
	"gonewild",
	"ass",
	"datgap",
	"pawg",
	"girlsinyogapants",
	// "rearpussy",
	"assinthong",
}
var perv_boys = []string{
	"HotGuys",
	"malemodels",
	"cuteguys",
	"hardbodiesmale",
	"malemodelsNSFW",
}
func get_reddit_one(ct *CT, reddit_perv []string) (string, error) {
	var a Reddit_Thread
	randy := ct.Random.Intn(len(reddit_perv))
	err := commonJsonRequest(ct, "https://www.reddit.com/r/"+reddit_perv[randy]+".json", &a)
	if err != nil {
		return "", err
	} else {
		return a.Data.Children[ct.Random.Intn(len(a.Data.Children))].Data.Preview.Images[0].Source.URL, nil
	}
}
func common_perv(ct *CT, s Slack.OMNI, reddit_perv []string) {
	if _,v:=ct.Slack.GetNameById(s.Channel);v == "general" {
		Message(ct.Websocket, s, "Pas dans le chan general, pervers !")
		return
	}
	Typing(ct.Websocket, s)
	img, err := get_reddit_one(ct, reddit_perv)
	if err!=nil{
		Message(ct.Websocket, s, err.Error())
		return
	}
	Typing(ct.Websocket, s)
	Message(ct.Websocket, s, img)
}
func perv(ct *CT, s Slack.OMNI) {
	common_perv(ct, s, perv_girls)
}
func perv_get_boys(ct *CT, s Slack.OMNI) {
	common_perv(ct, s, perv_boys)
}
// TODO(doc): reduce that ... thing
type Reddit_Thread struct {
	Data struct {
		Children []struct {
			Data struct {
				Domain        string        `json:"domain"`
				Subreddit     string        `json:"subreddit"`
				SelftextHTML  interface{}   `json:"selftext_html"`
				Selftext      string        `json:"selftext"`
				Likes         interface{}   `json:"likes"`
				SuggestedSort interface{}   `json:"suggested_sort"`
				UserReports   []interface{} `json:"user_reports"`
				SecureMedia   interface{}   `json:"secure_media"`
				LinkFlairText interface{}   `json:"link_flair_text"`
				ID            string        `json:"id"`
				FromKind      interface{}   `json:"from_kind"`
				Gilded        int           `json:"gilded"`
				ReportReasons interface{}   `json:"report_reasons"`
				Author        string        `json:"author"`
				Media         interface{}   `json:"media"`
				Score         int           `json:"score"`
				ApprovedBy    interface{}   `json:"approved_by"`
				Over18        bool          `json:"over_18"`
				Preview       struct {
					Images []struct {
						Source struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
							Nsfw struct {
								Source struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"source"`
								Resolutions []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"resolutions"`
							} `json:"nsfw"`
						} `json:"variants"`
						ID string `json:"id"`
					} `json:"images"`
				} `json:"preview"`
				NumComments      int    `json:"num_comments"`
				Thumbnail        string `json:"thumbnail"`
				SubredditID      string `json:"subreddit_id"`
				Downs            int    `json:"downs"`
				SecureMediaEmbed struct {
				} `json:"secure_media_embed"`
				PostHint        string        `json:"post_hint"`
				From            interface{}   `json:"from"`
				FromID          interface{}   `json:"from_id"`
				Permalink       string        `json:"permalink"`
				Name            string        `json:"name"`
				Created         float64       `json:"created"`
				URL             string        `json:"url"`
				AuthorFlairText interface{}   `json:"author_flair_text"`
				Title           string        `json:"title"`
				CreatedUtc      float64       `json:"created_utc"`
				Distinguished   interface{}   `json:"distinguished"`
				ModReports      []interface{} `json:"mod_reports"`
				NumReports      interface{}   `json:"num_reports"`
				Ups             int           `json:"ups"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
