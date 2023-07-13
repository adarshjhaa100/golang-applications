package webbasics

import (
	"fmt"
	"net/http"
)

func BasicHttpServer() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome to homepage!</h1>")
		fmt.Printf("Req object: %#v\n", r)
		
		// Read get parameters
		if method := r.Method; method == "GET" {
			fmt.Printf("Get req object parmas: %#v\n ", r.URL.Query().Get("a"))
		} else if method == "POST" {
			// Read post params (here html form value)
			fmt.Printf("POST req object parmas: %#v\n ", r.FormValue("test1"))
		}
	})

	// Http handler with contents of file system
	fs := http.FileServer(http.Dir("data/static/"))
	// http.Handle("/static/", fs)
	// domain/static/ will give access to tall the files. the /static/ is 
	// stripped before request reaches fileServer. Its because the fs has 
	// mounted data/static/ as root. so the contents indise the /static are 
	// visible to the FileServer. request to /static tells to find files in the
	// data/static/ + /static/ = data/static/static directory
	http.Handle("/static/", http.StripPrefix("/static/", fs)) 


	// Listens on TCP address, then calls Serve with handler
	const PORT = "8081"
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}
