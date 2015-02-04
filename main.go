package main

import (
	"fmt"
	"github.com/gorilla/mux"
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
	opts := &pg.Options{User: "joshrendek", Database: "cookingtips"}
	if os.Getenv("DB_HOST") != "" {
		opts = &pg.Options{
			User:     os.Getenv("DB_USER"),
			Database: "cookingtips",
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASS"),
		}
	}
	DB = pg.Connect(opts)
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

func Add(x, y int) int {
	return x + y
}

func Render(w http.ResponseWriter, name string, extra interface{}) {
	w.Header().Set("Content-Type", "text/html")
	tpl := template.New(name)
	tpl.Funcs(train.HelperFuncs)
	tpl.Funcs(template.FuncMap{
		"YoutubeURL": YoutubeURL,
		"add":        Add,
	})
	tpl, err := tpl.ParseFiles("templates/" + name + ".html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, name, extra)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "index", struct{ LandingImage string }{landingImage()})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "admin", nil)
}

func main() {
	fmt.Println("Setting up DB...")
	SetupDB()
	fmt.Println("Finished setting up DB...")
	train.ConfigureHttpHandler(nil)
	train.Config.BundleAssets = false
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/search", searchPageHandler)
	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/admin/pages", listPagesHandler)
	r.HandleFunc("/admin/page/{id:[0-9]+}", adminViewPageHandler)
	r.HandleFunc("/admin/pages/create", createPageHandler)
	r.HandleFunc("/tips/{id:[0-9]+}-{title}", viewPageHandler)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
