package main

import (
	// "fmt"
	// "net/http"
	"backend/internal/server"
	"log"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World! AC is tea time! 5555555")
// }

// func main() {
// 	http.HandleFunc("/", handler)

//		port := 8080
//		fmt.Printf("Server is running on http://localhost:%d\n", port)
//		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
//		if err != nil {
//			fmt.Printf("Error starting server: %s\n", err)
//		}
//	}
func main() {
	log.Println("Starting server...")
	server.StartServer()
}
