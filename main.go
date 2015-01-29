package main

import (
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

func main() {
	train.ConfigureHttpHandler(nil)
	train.Config.BundleAssets = false
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, nil)
}
