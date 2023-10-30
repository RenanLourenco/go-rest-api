package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-type","application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w,r)
	})
}