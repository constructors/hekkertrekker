package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

type PivotalTracker struct {
	Token string
	ProjectID int
}

type ptStories struct {
	Stories []ptStory `xml:"story"`
}

type ptStory struct {
	ID int `xml:"id"`
	Estimate int `xml:"estimate"`
	Name string `xml:"name"`
	Type string `xml:"story_type"`
	Labels string `xml:"labels"`
	CurrentState string `xml:"current_state"`
}

func (pt PivotalTracker) request(method string, path string) []byte {
	url := fmt.Sprintf("http://www.pivotaltracker.com/services/v3/projects/%d/%s",
	pt.ProjectID, path)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		bye("%v\n", err)
	}
	req.Header.Add("X-TrackerToken", pt.Token)
	if method == "PUT" {
		req.Header.Add("Content-Length", "0")
	}

	resp, err := client.Do(req)
	if err != nil {
		bye("%v\n", err)
	}
	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bye("%v\n", err)
	}

	return ret
}

func (pt PivotalTracker) AcceptableStories() []ptStory {
	var stories ptStories

	flags := url.Values{}
	flags.Add("filter", `state:unstarted,rejected type:Feature,Bug`)
	uri := fmt.Sprintf("stories?%s", flags.Encode())

	if err := xml.Unmarshal(pt.request("GET", uri), &stories); err != nil {
		bye("%v\n", err)
	}

	return stories.Stories
}

func (pt PivotalTracker) AcceptStory(id int) {
	flags := url.Values{}
	flags.Add("story[owned_by]", generalConfig.Name)
	flags.Add("story[current_state]", "started")
	uri := fmt.Sprintf("stories/%d?%s", id, flags.Encode())

	pt.request("PUT", uri)
}

func (pt PivotalTracker) DeliverStory(id int) {
	flags := url.Values{}
	flags.Add("story[current_state]", "delivered")
	uri := fmt.Sprintf("stories/%d?%s", id, flags.Encode())

	pt.request("PUT", uri)
}

func (pt PivotalTracker) DoneStory(id int) bool {
	var stories ptStories
	var labels = []string{}

	flags := url.Values{}
	flags.Add("filter", fmt.Sprintf("state:accepted id:%d", id))
	uri := fmt.Sprintf("stories?%s", flags.Encode())

	if err := xml.Unmarshal(pt.request("GET", uri), &stories); err != nil {
		bye("%v\n", err)
	}

	if (len(stories.Stories) == 0) {
		return false
	}

	if stories.Stories[0].Labels != "" {
		labels = strings.Split(stories.Stories[0].Labels, ",")
	}
	labels = append(labels, generalConfig.DoneLabel)

	flags = url.Values{}
	flags.Add("story[labels]",
		strings.Join(labels, ","))
	uri = fmt.Sprintf("stories/%d?%s", id, flags.Encode())

	pt.request("PUT", uri)

	return true
}

var pt PivotalTracker

func initPivotalTracker() {
	pt = PivotalTracker{generalConfig.Token, repositoryConfig.ProjectID}
}