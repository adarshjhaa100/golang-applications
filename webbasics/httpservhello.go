package webbasics

import (
	"fmt"
	// "mime/multipart"
	"net/http"
	"sync"
)

func checkErrHttp(err error){
	if err!=nil {
		panic(err)
	}
}


func helloWorldHttpServer(){
	wg := sync.WaitGroup{}
	
	// handler works on how to handle http client requests. 
	// first argument is URL pattern. Second is tun handler function
	
	// Register handler with a pattern
	http.HandleFunc("/", func(wrtr http.ResponseWriter, r *http.Request) {
		// w is a writer object that can be passed to a fprintf(writes to io writer) function
		fmt.Fprintf(wrtr, "Requested at : %#v", r.URL.Path)
		wg.Add(1)
		go func(){
			fmt.Printf("Requested: %#v\n\n", r)
			wg.Done()
		}()
		wg.Wait()
	})

	// listens on TCP network address and calls serve with the handler 
	const PORT = "8081"
	fmt.Println("listening on: ", PORT)
	
	http.ListenAndServe(fmt.Sprintf(":%s",PORT), nil)
	
}