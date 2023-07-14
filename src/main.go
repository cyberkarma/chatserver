package main

import (
	"github.com/cyberkarma/chatserver/cmd"
)

func main() {
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})
	//r.Get("/api/{greeter}", getGreeter)
	//http.ListenAndServe(":3000", r)

	cmd.Execute()
}

//func getGreeter(w http.ResponseWriter, r *http.Request) {
//	greeterParam := chi.URLParam(r, "greeter")
//	w.Write([]byte(greeterParam))
//
//}
