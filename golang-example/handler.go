package main

import "fmt"
import "net/http"


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}