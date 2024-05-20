package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates a slice of random birthdays represented as days of the year (1-365)
func generateRandomBirthdays(n int) []int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	birthdays := make([]int, n)      // Create a slice to hold the birthdays
	for i := range birthdays {
		birthdays[i] = rand.Intn(365) + 1 // Assign a random day (1-365) to each element
	}
	return birthdays
}

// Checks if there are any duplicate elements in the slice
func hasDuplicate(arr []int) bool {
	pigeonholes := make(map[int]bool) // Create a map to keep track of seen birthdays

	for _, num := range arr {
		if pigeonholes[num] { // If the birthday is already in the map, return true
			return true
		}
		pigeonholes[num] = true // Otherwise, add the birthday to the map
	}

	return false // If no duplicates are found, return false
}

// Runs multiple experiments to estimate the probability of at least one duplicate birthday
func calculateProbability(numExperiments int, numBirthdays int) float64 {
	count := 0 // Initialize count of experiments with at least one duplicate
	for i := 0; i < numExperiments; i++ {
		birthdays := generateRandomBirthdays(numBirthdays) // Generate a set of random birthdays
		if hasDuplicate(birthdays) {                       // Check if there is any duplicate
			count++ // Increment the count if a duplicate is found
		}
	}
	return float64(count) / float64(numExperiments) // Calculate the probability
}

func main() {
	numExperiments := 10000 // Number of experiments to run
	numBirthdays := 23      // Number of people (birthdays) in each experiment
	probability := calculateProbability(numExperiments, numBirthdays)
	fmt.Printf("After %d experiments, the probability of at least one duplicate birthday among %d people is %.2f\n",
		numExperiments, numBirthdays, probability)
}
