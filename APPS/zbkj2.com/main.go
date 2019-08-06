package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("start...\n")

	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServeTLS(":8080", "zbkj2.com.crt", "zbkj2.com.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Printf("end\n")
}
