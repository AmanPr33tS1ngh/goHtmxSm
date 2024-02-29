package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct{
	Title string
	Director string
}

func main()  {
	fmt.Println("go proj")

	fmt.Println("go proj")
	h1 := func(w http.ResponseWriter, r *http.Request)  {
		// io.WriteString(w, "Hello world!")
		// io.WriteString(w, r.Method)
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][] Film{
			"Films": {
				{
					Title: "Harry Potter", Director: "dunno",  
				},
				{
					Title: "Harry Potter2", Director: "dunno2",  
				},
				{
					Title: "Harry Potter3", Director: "dunno3",  
				},
			},
		}
		fmt.Print("/")
		tmpl.Execute(w, films)
	}
	h2 := func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("HTMX req received!")
		fmt.Println(r.Header.Get("HX-Request"))
		fmt.Println("h2")
		film := r.FormValue("title")
		director := r.FormValue("director")
		fmt.Println(film)
		fmt.Println(director)
		htmlStr := fmt.Sprintf("<li>%s - %s</li>", film, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add/", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
