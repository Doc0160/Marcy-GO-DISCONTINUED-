package Slack

import (
	"encoding/json"
	"net/http"
	URL "net/url"
	// "fmt"
)

type Slack struct {
	Token  string
	Client http.Client
	RTM    RTM
}
const(
	IM      = "im"
	USER    = "user"
	GROUP   = "group"
	CHANNEL = "channel"
)
// Call slack's web api
func (slack *Slack) API_CALL(method string, args map[string]interface{}) (interface{}, error) {
	url := "https://slack.com/api/" + method + "?token=" + slack.Token
	if method == "rtm.start" {
		resp, err := slack.Client.Get("https://slack.com/api/rtm.start?token=" + slack.Token)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(resp.Body).Decode(&slack.RTM)
		resp.Body.Close()
		return slack.RTM, err
	} else {
		for k, v := range args {
			if k == "attachments" {
				m, _ := json.Marshal(v)
				v = string(m)
			}
			url += "&" + k + "=" + URL.QueryEscape(v.(string))
		}
		resp, err := slack.Client.Get(url)
		if err != nil {
			return nil, err
		}
		var temp interface{}
		err = json.NewDecoder(resp.Body).Decode(&temp)
		resp.Body.Close()
		if err != nil {
			return temp, err
		}
		return temp, nil
	}
}

// return object of crorresponding name (either user,im,group,channel) in interface
func (slack *Slack) GetByName(name string, w string) (string, interface{}) {
	// w := slack.TypeOfId(id)
	//
	switch(w){
	case USER:
		for _, v := range slack.RTM.Users {
			if name == v.Name {
				return w, v
			}
		}
	case IM:
		for _, v := range slack.RTM.Ims {
			if name == v.User {
				return w, v
			}
		}
	case GROUP:
		for _, v := range slack.RTM.Groups {
			if name == v.Name {
				return w, v
			}
		}
	case CHANNEL:
		for _, v := range slack.RTM.Channels {
			if name == v.Name {
				return w, v
			}
		}
	}
	return "", nil
}
// return object of crorresponding id (either user,im,group,channel) in interface
func (slack *Slack) GetById(id string) (string, interface{}) {
	w := slack.TypeOfId(id)
	//
	switch(w){
	case USER:
		for _, v := range slack.RTM.Users {
			if id == v.ID {
				return w, v
			}
		}
	case IM:
		for _, v := range slack.RTM.Ims {
			if id == v.ID {
				return w, v
			}
		}
	case GROUP:
		for _, v := range slack.RTM.Groups {
			if id == v.ID {
				return w, v
			}
		}
	case CHANNEL:
		for _, v := range slack.RTM.Channels {
			if id == v.ID {
				return w, v
			}
		}
	}
	return "", nil
}
func (slack *Slack) GetNameById(id string) (string, string) {
	w, v := slack.GetById(id)
	//
	switch(w){
		case CHANNEL:
			return w, v.(Channel).Name
		case GROUP:
			return w, v.(Group).Name
		case USER:
			return w, v.(User).Name
		case IM:
			_, v = slack.GetNameById(v.(Im).User)
			return w, v.(string);
	}
	return "", ""
}

// (U)users, (D)ims, (G)groups, (C)channels
func (slack *Slack) TypeOfId(id string) string {
	id_types := map[byte]string{
		'U': USER,
		'D': IM,
		'G': GROUP,
		'C': CHANNEL,
	}
	if len(id) > 0 {
		if _, ok := id_types[id[0]]; ok {
			return id_types[id[0]]
		}
	}
	return ""
}
func(slack *Slack)SetPresence(id string, a string){
	for _, v := range slack.RTM.Users {
		if id == v.ID {
			v.Presence=a
		}
	}
}