package main

import (
	"fmt"
	"github.com/shaoshing/train"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

func landingImage() string {
	landingImages := []string{"blueberries.jpg", "coffee.jpg",
		"cutting_board.jpg", "ginger.jpg", "oranges.jpeg",
		"rasberries.jpg", "rasberries2.jpg", "redonion.jpg",
		"spices.jpg", "strawberries.jpg"}

	return landingImages[rand.Intn(len(landingImages))]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl := template.New("index")
	tpl.Funcs(train.HelperFuncs)
	tpl, err := tpl.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, "index", struct{ LandingImage string }{landingImage()})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl := template.New("admin")
	tpl.Funcs(train.HelperFuncs)
	tpl, err := tpl.ParseFiles("templates/admin.html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, "admin", nil)
}

type Page struct {
	Title        string
	Instructions string
	Youtubes     string
	Articles     string
	Tags         string
}

func createPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//_, _ := ioutil.ReadAll(r.Body)
	title := r.FormValue("title")
	instructions := r.FormValue("instructions")
	youtubes := r.FormValue("youtubes")
	articles := r.FormValue("articles")
	tags := r.FormValue("tags")
	page := Page{
		Title:        title,
		Instructions: instructions,
		Youtubes:     youtubes,
		Articles:     articles,
		Tags:         tags,
	}
	fmt.Println(page)
}

func main() {
	train.ConfigureHttpHandler(nil)
	train.Config.BundleAssets = false
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/admin/pages/create", createPageHandler)
	http.ListenAndServe(":"+port, nil)
}
