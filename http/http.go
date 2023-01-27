package main

import (
	"fmt"
	"net/http"
)

func Handler() func(w http.ResponseWriter, req *http.Request) {
	a := make([]string, 0)

	return func(w http.ResponseWriter, req *http.Request) {
		urlpath := req.URL.Path[1:]
		for i := range a {
			if a[i] == urlpath {
				b := "Welcome back" + urlpath
				fmt.Fprintln(w, b)
				return
			}
		}
		a = append(a, urlpath)
		b := "Greetings " + urlpath
		fmt.Fprintln(w, b)

	}
}

func main() {
	c := Handler()
	err := http.ListenAndServe(":8000", http.HandlerFunc(c))
	fmt.Println(err)
}
