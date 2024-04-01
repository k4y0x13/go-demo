package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// fileServer := http.FileServer(http.Dir("./static"))
	fmt.Printf("Sever listening at port 8080\n")

	// ? Create a home page handler
	http.HandleFunc("/", homePageHandler)
	// ? Create a hello handler
	http.HandleFunc("/hello", helloHandler)
	// ? About page handler
	http.HandleFunc("/about", aboutPageHander)
	// ? Login page hander
	http.HandleFunc("/login", loginPageHandler)

	// ? Start the server, if any error, console log it
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ? Handler for the home route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

// ? Handler for the about route
func aboutPageHander(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/about.html")
}

// ? Handler for the login route
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/login.html")
}

// ? Handler for the hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Page!")
}
