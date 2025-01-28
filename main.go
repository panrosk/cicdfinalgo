package main

import (
	"cicdfinalgo/utils"
	"fmt"
	"net/http"
)

func main() {
	who := utils.SayHiTo("Raymundo")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, who)
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
