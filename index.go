package main

import (
	"net/http"
	"./data"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)
		//public_tmpl_files := []string{
		//	"templates/layout.html",
		//	"templates/public.navbar.html",
		//	"templates/index.html"}
		//private_tmpl_files := []string{
		//	"templates/layout.html",
		//	"templates/private.navbar.html",
		//	"templates/index.html"}
		//var templates *template.Template
		if err != nil {
			//templates = template.Must(template.ParseFiles(private_tmpl_files...))
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			//templates = template.Must(template.ParseFiles(public_tmpl_files...))
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
		//templates.ExecuteTemplate(w, "layout", threads)
	}
}
