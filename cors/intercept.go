package cors

import (
	"log"
	"net/http"
)

//	Intercept is a middleware to
// 	inspect request for CORS header
//	adds necessary header to ResponseWriter if found.
//
func Intercept(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Println("about to intercept cors...")

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	next(rw, r)
}
