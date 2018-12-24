package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/handle2001/median"
)

type numberSet struct {
	Name string
	Set  []int
}

type response struct {
	Name  string
	Type  string
	Value []int
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d numberSet
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}
	m := median.GetMedian(d.Set)
	j := response{d.Name, "Median", m}

	s, err := json.Marshal(j)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintf(w, string(s))
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
