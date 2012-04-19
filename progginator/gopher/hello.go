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
	"fmt"
	"net/http"
)

// The init function is called when your application starts.
func init() {
	// Register a handle for /hello URLs.
	http.HandleFunc("/hello", hello)
}

// hello is an HTTP handler that prints "Hello Gopher!"
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gopher!")
}

// STOP OMIT
