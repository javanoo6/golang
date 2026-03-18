package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"chapters/ch7/datafile"
)

func main() {
	cwd_path, _ := os.Getwd()

	lines, err := datafile.GetStrings(cwd_path + "/ch7/datafile/votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

	// var ranks map[string]int
	// ranks = make(map[string]int)
	// ranks := make(map[string]int)

	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])

	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I haven't been assigned"])

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])

	// this would cause panic

	// var nilMap map[int]string
	// fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three"

	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)

	newcounters := map[string]int{"a": 3, "b": 0}
	var value int
	var contains bool
	value, contains = newcounters["a"]
	fmt.Println(value, contains)
	value, contains = newcounters["b"]
	fmt.Println(value, contains)
	_, contains = newcounters["c"]
	fmt.Println(contains)

	var ok bool
	newranks := make(map[string]int)
	var rank int
	newranks["bronze"] = 3
	rank, ok = newranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(newranks, "bronze")
	rank, ok = newranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	isPrime := make(map[int]bool)

	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}

	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}
