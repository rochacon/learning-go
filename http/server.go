package main

import (
	"io"
	"net/http"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	out := r.Method +" "+ r.URL.Path +"?"+ r.URL.RawQuery + "\n"
	print(out)
	io.WriteString(w, out)
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("Crashed: ", err)
	}
}
