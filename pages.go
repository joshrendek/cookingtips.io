package main

import (
	_ "gopkg.in/pg.v2"
	"net/http"
	"strings"
)

// create table pages(
//id serial primary key not null,
//title varchar(255),
//instructions text[],
//youtubes text[],
//articles text[],
//tags text[]
//);

type Page struct {
	Id           int64
	Title        string
	Instructions []string
	Youtubes     []string
	Articles     []string
	Tags         []string
}

func searchPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// select * from pages where title ~* 'dic';
}

func createPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	title := r.FormValue("title")
	instructions := strings.Split(r.FormValue("instructions"), "\n")
	youtubes := strings.Split(r.FormValue("youtubes"), "\n")
	articles := strings.Split(r.FormValue("articles"), "\n")
	tags := strings.Split(r.FormValue("tags"), "\n")
	page := Page{
		Title:        title,
		Instructions: instructions,
		Youtubes:     youtubes,
		Articles:     articles,
		Tags:         tags,
	}
	page.CreatePage()
}

func (p *Page) CreatePage() {
	_, err := DB.ExecOne(`INSERT INTO pages (title, instructions, youtubes, articles, tags) VALUES (?title, ?instructions, ?youtubes, ?articles, ?tags)`, p)
	if err != nil {
		panic(err)
	}
}
