package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// type User struct {
// 	Login string `json: "login"`
// 	ID    int    `json: "id"`

// 	FollowersURL string `json: "followers_url"`
// 	FollowingURL string `json: "following_url"`

// 	Type string `json: "type"`
// }

type Header struct {
	Key   string
	Value string
}

func main() {
	done := make(chan bool)
	go tick(done)

	time.Sleep(time.Second * 15)
}

func tick(done chan bool) {
	conf := GetConfig()

	interval := time.Duration(3600000 / conf.RateLimit)
	ticker := time.NewTicker(interval * time.Millisecond)

	// make savefile that can also generate

	var i int = 7000
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("Iteration", i, "at", t)
			i++

			m := []Header{
				{"Content-Length", "0"},
				{"Accept", "application/vnd.github.v3+json"},
				{"Authorization", fmt.Sprint("token ", conf.Token)},
			}

			// need to get username from endpoint /user/:ID and then send PUT request to endpoint /user/following/:USERNAME

			user := GetUser(i, conf)

			url := fmt.Sprint("https://api.github.com/user/following/", user.Login)

			_, err := call(url, "PUT", m, conf.HTTPTimeout)
			if err != nil {
				panic(err)
			}
		}
	}
}

func call(url, method string, headers []Header, timeout int) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("Error: %s", err.Error())
	}

	if headers != nil {
		for i := 0; i < len(headers); i++ {
			h := headers[i]

			req.Header.Add(h.Key, h.Value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error: %s", err.Error())
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(url, res.Status)

	return string(bytes), nil
}
