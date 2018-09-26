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
		reader := csv.NewReader(f)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

             for {
  		record, err := reader.Read()
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
