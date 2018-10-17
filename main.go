package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1> Header: Hello World</h1>
									<p> Paragraph: Go is great</p>
									<hr><p> Paragraph: and simple to learn</p>`)

	fmt.Fprintf(w, "<h1> Header: Hello World</h1>")
	fmt.Fprintf(w, "<p> Paragraph: Go is great</p>")
	fmt.Fprintf(w, "<hr><p> Paragraph: and simple to learn</p>")
	fmt.Fprintf(w, "<hr><p> You %s even add %s </p>", "can", "<strong>variables</strong>")

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)

}
