package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

type Article struct{
	//Content []string
	Id int
	Title string
	Content string
}

type Data struct{
	Content interface{}
}

func main(){
    http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))
	
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		info := []Article{
			{Id: 1, Title: "Artikel 1", Content: "Lorem Ipsum is simply dummy text"},
			{Id: 2, Title: "Artikel 2", Content: "Lorem Ipsum is simply dummy text"},
		};
		data := Data{Content: info};

		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
		))

		var err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"title": "Lorem Ipsum", "content": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}

		var tmpl = template.Must(template.ParseFiles(
			"views/content.html",
			"views/_header.html",
		))

		var err = tmpl.ExecuteTemplate(w, "content", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}