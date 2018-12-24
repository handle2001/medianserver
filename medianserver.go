package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
)

func getMedian(s []int) int {
	var median int
	// To find the median of a set takes  steps:
	// 1. Sort the set from smallest to largest
	sort.Ints(s)
	// 2. Find the length of the set
	slen := len(s)
	// 3. If the length of the set is odd, adding one to the length and then
	//		dividing that number by two gives the index of the median.
	if slen%2 != 0 {
		mIndex := slen / 2
		median = s[mIndex]

		// 4. If the length of the set is even, a few more steps are needed:
	} else if slen%2 == 0 {
		//	4a. Divide the length by two, this gives a float
		x := float64(slen / 2)
		//	4b. Round the float to the nearest integer up and down
		ceiling := int(math.Ceil(x))
		floor := int(math.Floor(x))
		//	4c. Average the values found at the two indexes found in 4b
		median = (s[ceiling] + s[floor]) / 2
	}

	return median
}

type numberSet struct {
	Name string
	Set  []int
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d numberSet
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "The median of your set called %s is %d\n", d.Name, getMedian(d.Set))
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
