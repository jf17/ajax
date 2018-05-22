package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type JSONData struct {
	Quote string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "ajax-json.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	//parse request to struct
	fmt.Println("Open /ajax")
	var d JSONData
	var answer JSONData

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(d.Quote)
	// create json response from struct

	answer.Quote = zipReader(d.Quote + ".ZIP")
	fmt.Println(answer.Quote)
	a, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(a)
	fmt.Println("Send response\n")
}

func zipReader(fileZIP string) string {
	razdelitel := "<br>" // потом поменять на <br> !!!

	// Open a zip archive for reading.

	r, err := zip.OpenReader(fileZIP)

	if err != nil {

		log.Fatal(err)

	}

	defer r.Close()

	// Iterate through the files in the archive,

	// printing some of their contents.

	result := razdelitel + razdelitel + "/home/user/archive/" + fileZIP + razdelitel

	for _, f := range r.File {

		name := " " + f.Name
		t := f.ModTime()
		last_time := t.Format("02-01-2006 15:04:05")
		last_string := "       " + last_time + razdelitel

		result = result + name + last_string

	}

	return result
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/ajax", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
