package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"net/http"
	"path/filepath"
	"strings"
)

type Post struct {
	Title string
	Date  string
	Body  template.HTML
}

var templates map[string]*template.Template

func main() {
	fmt.Println("Running stanley on port 8080...")

	loadTemplates()
	loadPosts()

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)

	post := r.PathPrefix("/post/{id}").Subrouter()
	post.Methods("GET").HandlerFunc(PostShowHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", r)
}

func NotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "404 not found")
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("posts")
	fileNames := make([]string, len(files))

	for _, f := range files {
		fileNames = append(fileNames, strings.TrimSuffix(f.Name(), ".md"))
	}

	tpl, ok := templates["listing"]
	if !ok {
		http.Error(rw, "Template not found", http.StatusInternalServerError)
		return
	}

	if err := tpl.ExecuteTemplate(rw, "base", fileNames); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	data, err := ioutil.ReadFile(fmt.Sprintf("parsed/%v.md.html", id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(rw, string(data))
}

func loadPosts() {
	files, _ := ioutil.ReadDir("posts/")

	for _, f := range files {
		data, err := ioutil.ReadFile(fmt.Sprintf("posts/%v", f.Name()))
		if err != nil {
			panic(err)
		}
		parsePost(f.Name(), data)
	}
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob("templates/*.html")
	if err != nil {
		panic(err)
	}

	for _, layout := range layouts {
		templateName := strings.TrimSuffix(filepath.Base(layout), ".html")

		if layout == "templates/base.html" {
			templates[templateName] = template.Must(template.ParseFiles(layout))
		} else {
			templates[templateName] = template.Must(template.ParseFiles(layout, "templates/base.html"))
		}
	}
}

func parsePost(id string, data []byte) {
	post := Post{}

	dataSplit := strings.Split(string(data), "<-->")
	config := dataSplit[0]
	err := yaml.Unmarshal([]byte(config), &post)
	if err != nil {
		fmt.Println(errors.New("parsePost: something went wrong parsing yaml."))
	}

	post.Body = template.HTML(blackfriday.MarkdownCommon([]byte(dataSplit[1])))

	tpl, ok := templates["post"]
	if !ok {
		fmt.Println("parsePost: template not found")
		return
	}

	f, err := os.Create(fmt.Sprintf("parsed/%v.html", id))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tpl.ExecuteTemplate(f, "base", post)
	if err != nil {
		panic(err)
	}

	f.Sync()
}
