package main

import (
	"fmt"
	"net/http"
)

func main() {

	HelloHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})
	HelloHandlerWithLog := logHandler(HelloHandler)

	http.Handle("/hello", HelloHandlerWithLog)
	http.ListenAndServe("0.0.0.0:8888", nil)

}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Before")
		h.ServeHTTP(writer, request)
		fmt.Println("After")
	})
}
