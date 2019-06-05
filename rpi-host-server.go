package main

import (
    "fmt"
//    "html/template"
    "net/http"

)
func hello (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to my website!")
}
func main() {    
	fmt.Println("starting")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)
}