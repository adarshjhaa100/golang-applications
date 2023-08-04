package webbasics

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func TemplateUse() {
	fmt.Println("Template Use!!!")

	// Parse the html file(s) as a dom tree
	templ := template.Must(
		template.ParseFiles("webbasics/templates/layout.html"))

	fmt.Printf("\nDefined Templates:\n %#v\n", templ.DefinedTemplates())
	// fmt.Printf("\nDefined Templates:\n %#v\n", templ.)

	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
		// Write the template to w
		templ.Execute(w, data)
	})

	mux.Handle("/static/",
	http.StripPrefix("/static/", http.FileServer(http.Dir("data/static/"))))

	
	http.ListenAndServe(":8081", mux)

}
