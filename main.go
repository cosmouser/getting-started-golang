package main

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
)

func cityHandler(res http.ResponseWriter, req *http.Request) {
	randomNumbers := [][]int{
		[]int{3214, 325342, 3432, 2323, 864732, 1256},
		[]int{85212, 13456, 416, 3267, 6783, 9824},
		[]int{52457, 3451, 1368, 89642, 7135, 476},
	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Amsterdam", "Berlin", "New York", "San Francisco", "Tokyo", "Santa Cruz"})
	table.AppendBulk(randomNumbers)
	table.Render()
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cities", cityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on port 5000 : ", err)
	}
}
