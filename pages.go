package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "gopkg.in/pg.v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//create table pages(
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
	Url          string // used to generate the url to go to
}

type Pages []*Page

func (pages *Pages) New() interface{} {
	p := &Page{}
	*pages = append(*pages, p)
	return p
}

func listPagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pages Pages
	_, err := DB.Query(&pages, "select * from pages")
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(pages)
	w.Write(js)
}

func viewPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	page := &Page{}
	_, err = DB.QueryOne(page, `SELECT * FROM pages WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}

	data := struct {
		Page         Page
		LandingImage string
	}{
		*page,
		landingImage(),
	}

	Render(w, "view", data)
}

func YoutubeURL(youtube string, x int) string {
	uri, _ := url.Parse(youtube)
	code := strings.Split(uri.RawQuery, "=")[1]
	autoplay := ""
	if x == 0 {
		autoplay = "?autoplay=1"
	}
	return fmt.Sprintf("http://www.youtube.com/embed/%s%s#%s", code, autoplay, uri.Fragment)
}

func (p *Page) Urlify() *Page {
	p.Url = strings.ToLower(fmt.Sprintf("/tips/%d-%s", p.Id, p.Title))
	return p
}

func searchPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pages Pages
	searchTerm := r.FormValue("q")
	_, err := DB.Query(&pages, "select * from pages where title ~* ?;", searchTerm)
	if err != nil {
		panic(err)
	}

	// TODO: maybe apply this to the struct of Pages instead
	for _, p := range pages {
		p.Urlify()
	}

	js, err := json.Marshal(pages)
	w.Write(js)
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
