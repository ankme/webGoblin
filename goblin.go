package main

import (
	"net/http"
//	"strings"
)

// simple page writer function
func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	//message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	
	w.Write([]byte(message))
}

// reply ping response
func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("../../../src")))
	http.HandleFunc("/ping", ping)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
