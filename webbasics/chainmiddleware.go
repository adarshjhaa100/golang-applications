package webbasics

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


type Middleware func(f http.HandlerFunc) http.HandlerFunc



func Logging() Middleware {

	mid := func(f http.HandlerFunc) http.HandlerFunc {
		
		return func (w http.ResponseWriter, r* http.Request) {
			
			start := time.Now()
			defer func(){
				elapsed := time.Since(start)
				log.Printf( "[%v] on %v, elapsed: [%v]", 
							r.Method, r.URL.Path, elapsed.String() )
			}()
			f(w, r)

		}
	}

	return mid

}

func Method(method string) Middleware {
	mid := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			if r.Method != method {
				// Throw http error. No further writes should be made to w
				// to get various http status code: http.Status...					
				http.Error(w, 
						   http.StatusText(http.StatusBadRequest), 
						   http.StatusBadRequest )
				
				// panic("pomp")				
				return
			}
			h(w, r)
		}
	}

	return mid
}


func defaultHandler( w http.ResponseWriter, r* http.Request) {
	fmt.Fprint(w, "Requested Default")
} 


/*
	given a handler h and middlewares m1,m2... in order, return handlers

	order: m2(m1(h))

	Mat Ryer's API building

*/

func ChainHandlers( h http.HandlerFunc, m ...Middleware ) http.HandlerFunc {

	for _, mid := range m {
		h = mid(h)
	}

	return h
}



func RunChainedMiddleware(){

	http.HandleFunc("/chained", 
		ChainHandlers( defaultHandler, Method("GET"), Logging() ) )

	http.ListenAndServe(":8081", nil)

}
