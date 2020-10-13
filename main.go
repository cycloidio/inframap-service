package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cycloidio/inframap/generate"
	"github.com/cycloidio/inframap/printer"
	"github.com/cycloidio/inframap/printer/factory"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read body request", http.StatusInternalServerError)
			return
		}

		g, _, err := generate.FromState(body, generate.Options{})
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to generate graph: %w", err), http.StatusInternalServerError)
			return
		}

		p, _ := factory.Get("dot")
		p.Print(g, printer.Options{}, w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
