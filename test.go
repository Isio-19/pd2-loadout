// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"net/http"
	"html/template"
	
	"log"
)

type TempStruct struct {
	Content string
}

func redirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "./index.html", http.StatusSeeOther)
}

func loadIndex(w http.ResponseWriter, r *http.Request) {
	// filename := "index.html"
	// body, err := os.ReadFile(filename)
	t, _ := template.ParseFiles("index.html")

	tempStruct := &TempStruct{Content: "Wolah le test"}

	t.Execute(w, tempStruct)

	//use a template to load the load
	// fmt.Fprintf(w, "<h1>This is a test</h1>")
}

func main() {
	http.HandleFunc("/", redirectToIndex)
	http.HandleFunc("/index.html", loadIndex)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
