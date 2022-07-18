package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          bool        `json:"hireable"`
	Bio               interface{} `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func GetUser(id int, conf *Config) *User {
	url := fmt.Sprint("https://api.github.com/user/", fmt.Sprint(id))

	h := []Header{
		{"Accept", "application/vnd.github.v3+json"},
		{"Authorization", fmt.Sprint("token ", conf.Token)},
	}

	resp, err := call(url, "GET", h, conf.HTTPTimeout)
	if err != nil {
		panic(err)
	}

	bytes := []byte(resp)
	result := &User{}

	if err := json.Unmarshal(bytes, result); err != nil {
		panic(err)
	}

	return result
}
