package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!  THIS WAS DONE FROM THE GITHUB SITE</h1>")
}

// Added some comments

// Added some more comments

// Added yet more comments to check for syncing

// Comments added on the macbook pro 13

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
