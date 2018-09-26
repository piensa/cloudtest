package main

import (
	"fmt"
	"net/http"
	"encoding/csv"
  	"log"
  	"io"
        "os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
               fileName := "data.csv"
               f, err := os.Open(fileName)
		r := csv.NewReader(f)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

             for {
  		record, err := r.Read()
  		if err == io.EOF {
  			break
  		}
  		if err != nil {
  			log.Fatal(err)
  		}
  
  		fmt.Println(record)
      	}
	})

	http.ListenAndServe(":80", nil)
}
