package main

import (
	"fmt"
	"io/ioutil"
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
	resp, _ := http.Get("https://www.coindesk.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()

}
