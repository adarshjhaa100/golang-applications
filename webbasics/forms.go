package webbasics

import (
	"fmt"
	"html/template"
	"net/http"
	_ "time"
)

type Message struct {
	Success bool;
}

type FormBody struct {
	name string;
	phone string;
	dateTime string;
}

func HttpForms(){
	fmt.Printf("\nHTTP Forms\n")

	mtx := http.NewServeMux()
	mtx.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		
		tmpl := template.Must(
			template.ParseFiles("webbasics/templates/form.html"))
	
		data := Message {
			Success: false,
		}

		if(r.Method == "GET"){
			tmpl.Execute(w, data)
		}

		if( r.Method == "POST" ){
			formData := FormBody {
				name : r.FormValue("name"),
				phone: r.FormValue("phone"),
				dateTime: r.FormValue("dob"),
			}

			fmt.Printf("Form Data: \n %#v", formData)

			tmpl.Execute(w, Message {
				Success: true,
			})

		}
		
})
	http.ListenAndServe(":8081", mtx)
}