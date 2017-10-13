package basics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

/*
For Challenge 4:
Detect single-character XOR
One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.
*/

func ParallelXorSolve(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	bestAttemptChan := make(chan *Attempt)

	solve := func(cypher string) {
		defer wg.Done()
		results := BruteForceXorCrack(cypher)
		bestAttemptChan <- results[0]
	}

	count := 0
	for scanner.Scan() {
		cypher := scanner.Text()
		wg.Add(1)
		go solve(cypher)
		count++
	}

	// Close channel when the workers are finished
	go func() {
		wg.Wait()
		close(bestAttemptChan)
	}()

	bestAttempts := make([]Attempt, count)
	for a := range bestAttemptChan {
		bestAttempts = append(bestAttempts, *a)
	}
	sort.Slice(bestAttempts, func(i, j int) bool { return bestAttempts[i].Score > bestAttempts[j].Score })

	fmt.Printf("%0f: %s\n", bestAttempts[0].Score, bestAttempts[0].Text)
	return bestAttempts[0].Text
}
