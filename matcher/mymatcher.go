package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type MyMatcherPlugin struct {
	Count int
}

// Load loads the plugin instance
func Load() interface{} {
	return &MyMatcherPlugin{}
}

func (i *MyMatcherPlugin) MatcherFunc (req *http.Request) bool {
	i.Count += 1

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can't read body")
		return false
	}

	defer func() {
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}()

	log.Printf(">>> [%d] body: %s", i.Count, body)

	if string(body) == "hello" {
		return true
	}

	return false
}
