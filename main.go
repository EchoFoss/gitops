package main

import "net/http"

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/", SayHello)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Hello World!</h1>"))
}
