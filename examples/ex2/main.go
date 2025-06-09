package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
	"text/template"

	"github.com/nicolito128/gama"
)

var templates = template.Must(template.ParseFiles("./index.html"))

func display(w http.ResponseWriter, page string, data any) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "index", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func main() {
	http.HandleFunc("/", uploadHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error Retrieving the File: %v\n", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding the file: %v\n", err)
		return
	}

	pl := gama.New(img)

	colors, err := pl.Quantify(16)
	if err != nil {
		fmt.Printf("Error getting the color palette: %v\n", err)
		return
	}

	hexs := make([]string, len(colors))
	for i, c := range colors {
		hexs[i] = fmt.Sprintf("<div style=\"background: %s; width: 50px; height: 50px; display: inline-block;\"></div>", gama.ColorToHex(c, false))
	}

	w.Header().Set("Content-Type", "text/html")
	s := strings.Join(hexs, " ")
	display(w, "index", map[string]any{
		"Colors": s,
	})
}
