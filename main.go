package main

import (
	"github.com/shaoshing/train"
	"gopkg.in/pg.v2"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

var err error
var DB *pg.DB

func SetupDB() {
	DB = pg.Connect(&pg.Options{User: "joshrendek", Database: "cookingtips"})
	if err != nil {
		panic(err)
	}
}

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

func main() {
	SetupDB()
	train.ConfigureHttpHandler(nil)
	train.Config.BundleAssets = false
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchPageHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/admin/pages/create", createPageHandler)
	http.ListenAndServe(":"+port, nil)
}
