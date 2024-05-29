package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}
	err=ts.Execute(w,nil)
		if err !nil {
			log.Println(err.Error())
		}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Display the selected note ID @d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to our snippetbox"))
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, " Get-Method !", 405)
		return
	}
	w.Write([]byte("Creating a new note..."))
}
