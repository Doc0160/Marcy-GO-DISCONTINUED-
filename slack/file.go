package Slack
import(

)

type(

File struct{
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Title              string  `json:"title"`
	Size               float64 `json:"size"`
	Mimetype           string  `json:"mimetype"`
	Filetype           string  `json:"filetype"`
	PrettyType         string  `json:"pretty_type"`
	User               string  `json:"user"`
	Permalink          string  `json:"permalink"`
	Permalink_Public   string  `json:"permalink_public"`
	URLPrivate         string  `json:"url_private"`
	URLPrivateDownload string  `json:"url_private_download"`
}

)
