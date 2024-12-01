package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// Extract the lists from the file
	list1, list2 := getLists()

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	distSum := calculateSum(list1, list2)
	log.Printf("Distance Sum: %d\n", distSum)

	simScore := calculateSimilarityScore(list2, list1)
	log.Printf("Similarity score: %d\n", simScore)
}

// Calculate the similarity score (task 2)
func calculateSimilarityScore(list2 []int, list1 []int) int {
	// Create map of second list for quicker lookup
	l2Map := make(map[int]int, len(list2))
	for _, v := range list2 {
		if val, exists := l2Map[v]; !exists {
			l2Map[v] = 1
		} else {
			l2Map[v] = val + 1
		}
	}

	// Now, go over list1 again and multiply entries by map values from list2
	simScore := 0
	for _, v := range list1 {
		if multiplier, exists := l2Map[v]; exists {
			simScore = simScore + v*multiplier
		}
	}
	return simScore
}

// Calculate the distance (task 1)
func calculateSum(list1 []int, list2 []int) int {
	distSum := 0
	for i, v := range list1 {
		dist := list2[i] - v
		if dist <= 0 {
			dist = -dist
		}
		distSum = distSum + dist
	}
	return distSum
}

// Extract integer numbers from input file (predefined format, just 2 ints per
// line with whitespace in between)
func getLists() ([]int, []int) {
	file, err := os.Open("./01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var val1, val2 int
		if n, err := fmt.Sscan(scanner.Text(), &val1, &val2); n != 2 || err != nil {
			log.Fatalf("incorrect input structure, n=%d, err=%v", n, err)
		}
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list1, list2
}
