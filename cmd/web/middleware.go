package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// //WriteToConsole is middleware function
// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(w, r)
// 	})

// }

func NoSurf(next http.Handler) http.Handler {
	csrtHandler := nosurf.New(next)

	csrtHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrtHandler
}
