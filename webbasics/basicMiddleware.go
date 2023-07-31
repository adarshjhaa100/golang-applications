package webbasics

import (
	"math/rand"
	"fmt"
	"log"
	"net/http"
)

/*
	Middleware is something which is in the middle of 2 componentsand enables
	interaction between them. e.g.
	1. Between user and Backend system: API
	2. Between user and kernel: shell

	In go http, middleware takes a handlerFunc and returns another handlerFunc
*/

/*
	http.Handler is an interface with the method signature
	ServeHttp(Writer, *Request)

	handlerFunc is a function which implements ServeHttp

*/

func check(e error){
	if e != nil {
		panic(e)
	} 
}

/*
	The HandlerFunc type is an adapter to allow the use of ordinary functions 
	as HTTP handlers. 
	If f is a function with the appropriate signature, 
	HandlerFunc(f) is a Handler that calls f
*/

func basicMiddlewareLogging( f http.HandlerFunc ) http.HandlerFunc {
	
	return func( w http.ResponseWriter, r* http.Request ){
		// Write to logger
		log.Printf( "[%v] on %v", r.Method, r.URL.Path )
		f(w,r)		
	}

}


func testHandlerFunc(w http.ResponseWriter, r* http.Request) {	
	// Write to response writer. Client will see this
	_, err := fmt.Fprintf( w, fmt.Sprintf("Test: %v", rand.Intn(1000000)) )
	check(err)

}

func useBasicMiddleware(){
	
	http.HandleFunc("/path1", basicMiddlewareLogging(testHandlerFunc))
	http.HandleFunc("/path2", basicMiddlewareLogging(testHandlerFunc))
	
	http.ListenAndServe(":8081", nil)

}

func RunMiddlewares(){
	useBasicMiddleware()

}
