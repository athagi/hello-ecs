package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hash := os.Getenv("COMMIT_HASH")
	message := "Hello world \n" + hash
	fmt.Fprint(w, message)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
