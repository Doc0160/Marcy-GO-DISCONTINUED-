package Slack

type(

File struct {
	ID string `json:"id"`
	Created int `json:"created"`
	Timestamp int `json:"timestamp"`
	Name string `json:"name"`
	Title string `json:"title"`
	Mimetype string `json:"mimetype"`
	Filetype string `json:"filetype"`
	PrettyType string `json:"pretty_type"`
	User string `json:"user"`
	Mode string `json:"mode"`
	Editable bool `json:"editable"`
	IsExternal bool `json:"is_external"`
	ExternalType string `json:"external_type"`
	Username string `json:"username"`
	Size int `json:"size"`
	URLPrivate string `json:"url_private"`
	URLPrivateDownload string `json:"url_private_download"`
	Thumb64 string `json:"thumb_64"`
	Thumb80 string `json:"thumb_80"`
	Thumb360 string `json:"thumb_360"`
	Thumb360Gif string `json:"thumb_360_gif"`
	Thumb360W int `json:"thumb_360_w"`
	Thumb360H int `json:"thumb_360_h"`
	Thumb480 string `json:"thumb_480"`
	Thumb480W int `json:"thumb_480_w"`
	Thumb480H int `json:"thumb_480_h"`
	Thumb160 string `json:"thumb_160"`
	Permalink string `json:"permalink"`
	PermalinkPublic string `json:"permalink_public"`
	EditLink string `json:"edit_link"`
	Preview string `json:"preview"`
	PreviewHighlight string `json:"preview_highlight"`
	Lines int `json:"lines"`
	LinesMore int `json:"lines_more"`
	IsPublic bool `json:"is_public"`
	PublicURLShared bool `json:"public_url_shared"`
	DisplayAsBot bool `json:"display_as_bot"`
	Channels []string `json:"channels"`
	Groups []string `json:"groups"`
	Ims []string `json:"ims"`
	InitialComment struct {
	} `json:"initial_comment"`
	NumStars int `json:"num_stars"`
	IsStarred bool `json:"is_starred"`
	PinnedTo []string `json:"pinned_to"`
	Reactions []struct {
		Name string `json:"name"`
		Count int `json:"count"`
		Users []string `json:"users"`
	} `json:"reactions"`
	CommentsCount int `json:"comments_count"`
}
)
