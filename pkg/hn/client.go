package hn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MyClient interface {
	Maxitem() (int, error)
	Get(id int) (Item, error)
}

type Item struct {
	Id     int    `json:"id"`
	Author string `json:"by"`
	Title  string `json:"title"`
	Score  int    `json:"score"`
	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Type   string `json:"type"`
	Kids   []int  `json:"kids"`
}

type HttpClient struct {
	url string
}

const baseURL = "https://hacker-news.firebaseio.com/v0"

func NewHTTPClient() *HttpClient {
	return &HttpClient{baseURL}
}

func NewHTTPClientFor(url string) *HttpClient {
	return &HttpClient{url}
}

func (hc *HttpClient) Maxitem() (int, error) {
	resp, err := http.Get(hc.url + "/maxitem.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(body))
}

func (hc *HttpClient) Get(id int) (Item, error) {
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", hc.url, id))
	if err != nil {
		return Item{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Item{}, err
	}
	var item Item
	return item, json.Unmarshal(body, &item)
}
