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
	"fmt"
	"net/http"
	"time"
)

// The init function is called when your application starts.
func init() {
	// Register a handle for /progginator URLs.
	http.HandleFunc("/progginator", progginator)
	http.HandleFunc("/progginator_", progginator_)
}

func progginator(w http.ResponseWriter, r *http.Request) {
	begin := time.Now() // OMIT
	c := appengine.NewContext(r)
	hnItems := hackernewsItems(c)
	pgItems := proggitItems(c)
	w.Header().Add("Content-Type", "text/plain")
	for _, item := range hnItems {
		fmt.Fprintf(w, "%s: %s\n", item.Title, item.Url)
	}
	for _, item := range pgItems {
		fmt.Fprintf(w, "%s: %s\n", item.Title, item.Url)
	}
	c.Infof("progginator: %s", time.Since(begin)) // OMIT
}

// STOP OMIT

func progginator_(w http.ResponseWriter, r *http.Request) {
	begin := time.Now() // OMIT
	c := appengine.NewContext(r)
	results := make(chan []Item)
	go func() {
		results <- hackernewsItems(c)
	}()
	go func() {
		results <- proggitItems(c)
	}()
	w.Header().Add("Content-Type", "text/plain")
	for i := 0; i < 2; i++ {
		items := <-results
		for _, item := range items {
			fmt.Fprintf(w, "%s: %s\n", item.Title, item.Url)
		}
	}
	c.Infof("progginator_: %s", time.Since(begin)) // OMIT
}

// STOP OMIT
