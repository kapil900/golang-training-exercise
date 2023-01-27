package main

import (
	"fmt"
	"net/http"
)

func Handler() func(w http.ResponseWriter, req *http.Request) {
	//a := make([]string, 0)
	a := map[string]string{}

	return func(w http.ResponseWriter, req *http.Request) {
		urlpath := req.URL.Path[1:]
		for b, _ := range a {
			if a[b] == urlpath {
				d := "Welcome back" + urlpath
				fmt.Fprintln(w, d)
				return
			}
		}
		// for i := range a {
		// 	if a[i] == urlpath {
		// 		b := "Welcome back" + urlpath
		// 		fmt.Fprintln(w, b)
		// 		return
		// 	}
		// }
		//  c:= map[l]{
		// 	"a":1,

		//  }
		//  _,exist:=[c]
		// a = append(b, urlpath)
		b := "Greetings " + urlpath
		fmt.Fprintln(w, b)

	}
}

func main() {
	c := Handler()
	err := http.ListenAndServe(":8000", http.HandlerFunc(c))
	fmt.Println(err)
}
