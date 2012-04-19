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

const ProggitFeedUrl = "http://reddit.com/r/programming/.json"

type RedditFeed struct {
	Data RedditFeedData
}

type RedditFeedData struct {
	Items []RedditFeedItem `json:"children"`
}

type RedditFeedItem struct {
	Data Item
}

// STOP OMIT

const HNFeedUrl = "http://news.ycombinator.com/rss"

type HNFeed struct {
	Data HNFeedData `xml:"channel"`
}

type HNFeedData struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Url   string `xml:"link"`
}

// STOP OMIT



