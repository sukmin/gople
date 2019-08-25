package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const PORT  = 8080

func main() {
	portAddress := ":" + strconv.Itoa(PORT)

	http.HandleFunc("/", func (w http.ResponseWriter,r *http.Request) {
		fmt.Fprintln(w, "<!DOCTYPE html>")
		fmt.Fprintln(w, "<html lang=\"ko\">")
		fmt.Fprintln(w, "<head>")
		fmt.Fprintln(w, "<meta charset=\"utf-8\">")
		fmt.Fprintln(w, "<title>권석민의 블로그</title>")
		fmt.Fprintln(w, "</head>")
		fmt.Fprintln(w, "<body>")
		fmt.Fprintln(w, "hello world")
		fmt.Fprintln(w, "</body>")
		fmt.Fprint(w, "</html>")
	})

	http.ListenAndServe(portAddress,nil)
}
