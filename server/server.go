package server

import (
	"ascii-art-web-dockerize/asciiart"
	"ascii-art-web-dockerize/hashaddress"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Struct to hold the error data
type ErrorPageData struct {
	Code     string
	ErrorMsg string
}

// Struct to hold the result data
type ResultPageData struct {
	Input  string
	Banner string
	Result string
}

// Function to render the error page
func errHandler(w http.ResponseWriter, r *http.Request, err *ErrorPageData) {
	_ = r
	errorTemp := template.Must(template.ParseFiles("templates/error.html"))
	errorTemp.Execute(w, err)
}

// Function to render the main page
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Validating the request path
	if r.URL.Path != "/" {
		err := ErrorPageData{Code: "404", ErrorMsg: "PAGE NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errHandler(w, r, &err)
		return
	}
	// Validating the request method
	if r.Method != "GET" {
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		errHandler(w, r, &err)
		return
	}
	// Validating the parsing of the main page
	main, err := template.ParseFiles("templates/index.html")
	if err != nil {
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errHandler(w, r, &err)
		return
	}

	mainTemp := template.Must(main, nil)
	mainTemp.Execute(w, nil)
}

// Function to render the result page
func ResultHandler(w http.ResponseWriter, r *http.Request) {

	// Validating the request method
	if r.Method != "POST" {
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		errHandler(w, r, &err)
		return
	}

	// Validating the parsing of the form
	if err := r.ParseForm(); err != nil {
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errHandler(w, r, &err)
		return
	}

	// Validation for the input
	input := r.PostFormValue("input-text")
	inputValidation := strings.ReplaceAll(input, "\r\n", "")

	for _, letter := range inputValidation {
		if letter < 32 || letter > 126 {
			err := ErrorPageData{Code: "400", ErrorMsg: "INVALID INPUT"}
			w.WriteHeader(http.StatusBadRequest)
			errHandler(w, r, &err)
			return
		}
	}
	// Validation for the banner
	banner := r.PostFormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		err := ErrorPageData{Code: "404", ErrorMsg: "BANNER NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errHandler(w, r, &err)
		return
	}

	// Check if the banner file's hash is valid
	bannerFile := "banners/" + banner + ".txt"

	bannerHashes := map[string]string{
		"standard":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"shadow":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"thinkertoy": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
	}
	expectedHash, exists := bannerHashes[banner]
	if !exists {
		err := ErrorPageData{Code: "404", ErrorMsg: "BANNER NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errHandler(w, r, &err)
		return
	}

	err := hashaddress.ValidateFileHash(bannerFile, expectedHash)
	if err != nil {
		log.Printf("Error: Banner file hash mismatch for %s: %v\n", banner, err)
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errHandler(w, r, &err)
		return
	}

	// Use SplitNewLine function to process input if needed
	lines := asciiart.SplitNewLine([]string{input})
	log.Println("Processed lines: ", lines)
	// Validation for ASCII art generation
	ascii, err := asciiart.AsciiArt(strings.Join(lines, "\n"), banner)
	if err != nil {
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errHandler(w, r, &err)
		return
	}

	resultTemp := template.Must(template.ParseFiles("templates/layout.html"))
	//log.Println(ascii)

	output := ResultPageData{Input: input, Banner: banner, Result: ascii}

	resultTemp.Execute(w, output)
}
