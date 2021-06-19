package main

import (
	"fmt"
	"net/http"
)

//WriteToConsole is middleware function
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}
