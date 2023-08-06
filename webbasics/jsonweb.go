package webbasics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func JSONWeb() {
	mux := http.NewServeMux()

	mux.HandleFunc("/encode",
		func(w http.ResponseWriter, r *http.Request) {
			enc := json.NewEncoder(w)
			w.Header().Set("Content-Type", "application/json")
			enc.Encode(
				User{
					Name:      "Sam",
					Email:     "sam@gmail.com",
					Age:       12,
					Timestamp: time.Now(),
					Options:   []string{"dsiuhfhids", "dsfds"},
				},
			)

		})

	mux.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			var v any
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&v)
			check(err)

			// logger := log.Logger{}
			fmt.Printf("\nRequest body: %#v\n", v)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
		}

		w.WriteHeader(405)
	})

	http.ListenAndServe(":8082", mux)

}
