// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

//go:generate rice embed-go
var box *rice.Box

func main() {
	box = rice.MustFindBox("assets")
	log.Printf("Running website on http://127.0.0.1.xip.io:8080")
	http.Handle("/password.txt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get("X-Udacity-Exercise") != "" {
			w.Write([]byte("Password: piquizahhai5aeh2fah9Uk"))
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	}))
	http.Handle("/", http.FileServer(box.HTTPBox()))
	http.ListenAndServe(":8080", nil)
}
