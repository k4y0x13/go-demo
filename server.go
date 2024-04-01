package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	// fileServer := http.FileServer(http.Dir("./static"))
	fmt.Printf("Sever listening at port 10000\n")

	// ? Create a home page handler
	http.HandleFunc("/", homePageHandler)
	// ? Create a hello handler
	http.HandleFunc("/hello", helloHandler)
	// ? About page handler
	http.HandleFunc("/about", aboutPageHander)
	// ? Login page hander
	http.HandleFunc("/login", loginPageHandler)
	// ? auth resource hander
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/restricted", restrictedPageHandler)

	// ? Start the server, if any error, console log it
	if err := http.ListenAndServe("0.0.0.0:1000", nil); err != nil {
		log.Fatal(err)
	}
}

// ? Authentication handler
func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============== Authentication Route Called ============")
	// ? Get the Body of the request
	body, _ := io.ReadAll(r.Body)
	bodyString := string(body)

	// ? Create a JSON Object based on the request JSON String
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(bodyString), &jsonMap)
	// ? Printing to the console so that i can see the data
	fmt.Println("JSON: ", jsonMap)
	fmt.Println("Username: ", jsonMap["username"])
	fmt.Println("Password: ", jsonMap["password"])

	// ? Get the username and password from the JSON Object and put it to a variable of type string
	username := jsonMap["username"].(string)
	password := jsonMap["password"].(string)

	// ? Check if the username and password are correct
	if username == "admin" && password == "admin" {
		// ? If correct, send a json response with a success message
		jsonResponse := `{"message": "Login Successful!"}`

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, jsonResponse)

	} else {
		// ? If incorrect, send an error json message
		jsonResponse := `{"message": "Login Un-Successful!"}`

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, jsonResponse)
	}
}

// ? Handler for the home route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

// ? Handler for the restricted route
func restrictedPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/restricted.html")
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
