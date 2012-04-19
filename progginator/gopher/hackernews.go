// Copyright 2012 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gopher

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func init() {
	// Register a handler for /hackernews URL.
	http.HandleFunc("/hackernews", hackernews)
}

// STOP OMIT

func hackernews(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	w.Header().Add("Content-Type", "text/plain")
	for _, item := range hackernewsItems(c) {
		fmt.Fprintf(w, "%s: %s\n", item.Title, item.Url)
	}
}

// STOP OMIT

func hackernewsItems(c appengine.Context) []Item {
	begin := time.Now() // OMIT
	client := urlfetch.Client(c)
	resp, err := client.Get(HNFeedUrl)
	if err != nil {
		c.Errorf("Error fetching %s: %s", HNFeedUrl, err.Error())
		return []Item{}
	}
	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	var feed *HNFeed
	if err = decoder.Decode(&feed); err != nil {
		c.Errorf("Error decoding HNFeed: %s", err.Error())
		return []Item{}
	}
	c.Infof("hackernews: %s", time.Since(begin)) // OMIT
	return feed.Data.Items
}

// STOP OMIT
