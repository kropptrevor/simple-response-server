package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func main() {
	err := http.ListenAndServe(":3333", &handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	panic("something went wrong")
}
