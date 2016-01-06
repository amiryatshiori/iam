package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):]
	//p, _ := loadPage(title)
	p := &Page{Title: "Title1", Body: []byte("Body1")}
	renderTemplate(w, "view", p)
}

func main() {
    ip := os.Getenv("IP")
    port := os.Getenv("PORT")
    fmt.Printf("ip: %v %v\n", ip, port)
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":"+port, nil)
}