package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

/*
	Steps To Perform Sieve Of Eratosthenes

	1 - Generate numbers from 2 to T
	2 - Traverse from the smallest prime number which is num = 2
	3 - Eliminate or mark all the multiples of num which are lesser than or equal to T.
			It will help remove non-prime numbers and will help to reduce our number of comparisons to check
	4 - Update the value of num to the mediate next prime number. The next prime is the next number in the list
			which is greater than the current number
	5 - Repeat step three until num <= √T
*/
func main() {
	v1 := time.Now()
	getAllPrimes(15)
	log.Printf("main, execution time %s\n", time.Since(v1))

	v2 := time.Now()
	getAllPrimesV2(15)
	log.Printf("main, execution time %s\n\n", time.Since(v2))

	v3 := time.Now()
	fmt.Println(SieveOfEratosthenes(15))
	log.Printf("main, execution time %s\n", time.Since(v3))

	v1 = time.Now()
	getAllPrimes(127)
	log.Printf("main, 127 v1 execution time %s\n", time.Since(v1))

	v2 = time.Now()
	getAllPrimesV2(127)
	log.Printf("main, 127 v2 execution time %s\n", time.Since(v2))

	v3 = time.Now()
	SieveOfEratosthenes(127)
	log.Printf("main, 127 v3 execution time %s\n", time.Since(v3))

	v1 = time.Now()
	getAllPrimes(250)
	log.Printf("main, 250 v1 execution time %s\n", time.Since(v1))

	v2 = time.Now()
	getAllPrimesV2(250)
	log.Printf("main, 250 v2 execution time %s\n", time.Since(v2))

	v3 = time.Now()
	SieveOfEratosthenes(250)
	log.Printf("main, 250 v3 execution time %s\n", time.Since(v3))

	v1 = time.Now()
	getAllPrimes(599)
	log.Printf("main, 599 v1 execution time %s\n", time.Since(v1))

	v2 = time.Now()
	getAllPrimesV2(599)
	log.Printf("main, 599 v2 execution time %s\n", time.Since(v2))

	v3 = time.Now()
	SieveOfEratosthenes(599)
	log.Printf("main, 599 v3 execution time %s\n", time.Since(v3))

	//fmt.Println(getAllPrimes(127))
	//fmt.Println(getAllPrimes(250))
	//fmt.Println(getAllPrimes(599))
}

func getAllPrimes(limit int) []int {
	//Check for illegal limit
	if limit < 2 {
		return nil
	}

	// Creation of the Number Array Step 1
	var numberArray []int
	for i := 2; i <= limit; i++ {
		numberArray = append(numberArray, i)
	}

	primeNum := float64(numberArray[0])
	generalIndex := 0

	// Check in step 5
	for primeNum <= math.Sqrt(float64(limit)) {

		// Always starting the iteration over the array from the letter index that last prime was on if the last prime t
		//hat we have found is 3, and it is on index 2 we will iterate over the array [4...n]

		index := generalIndex + 1

		// Iteration of the array  Step 2
		for index < len(numberArray) {
			if numberArray[index]%int(primeNum) == 0 {
				// Elimination of the element if it is divisible by our prime number Step 3
				numberArray = remove(numberArray, index)
				index--
			}
			index++
		}
		// Updating the value of the next prime Step 5
		generalIndex++
		primeNum = float64(numberArray[generalIndex])
	}
	return numberArray
}

func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}
	var rtrArray []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			rtrArray = append(rtrArray, p)
		}
	}
	return rtrArray
}

func getAllPrimesV2(limit int) []int {
	//Check for illegal limit
	if limit < 2 {
		return nil
	}

	// Creation of the Number Array Step 1
	var numberArray []int
	for i := 2; i <= limit; i++ {
		numberArray = append(numberArray, i)
	}

	primeIndex := 0
	primeNum := numberArray[primeIndex]

	// Check in step 5
	for primeNum*primeNum <= limit {

		// Always starting the iteration over the array from the letter index that last prime was on if the last prime t
		//hat we have found is 3, and it is on index 2 we will iterate over the array [4...n]

		index := primeIndex + 1

		// Iteration of the array  Step 2
		for index < len(numberArray) {
			if numberArray[index]%primeNum == 0 {
				// Elimination of the element if it is divisible by our prime number Step 3
				numberArray[index] = -1
			} else {
				index++
			}

		}
		// Updating the value of the next prime Step 5
		primeIndex++
		primeNum = numberArray[primeIndex]
		if primeNum == -1 {
			for primeNum == -1 {
				primeIndex++
				primeNum = numberArray[primeIndex]
			}
		}

	}
	var resultArray []int
	for _, v := range numberArray {
		if v != -1 {
			resultArray = append(resultArray, v)
		}
	}
	return resultArray
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
