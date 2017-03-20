package main

import "fmt"

const limit = 1000

func main() {

	for _, prime := range Sieve() {
		fmt.Println(prime)
	}

}

//Sieve of Eratosthenes
func Sieve() []uint {
	var list = make([]bool, limit, limit)
	initialize(list)

	var prime uint = 2

	for ; prime < limit; {
		mark_multiples(prime, list)
		prime = get_next_prime(prime, list)
	}

	return get_unmarked_numbers(list)
}

func mark_multiples(prime uint, list []bool) {
	var multiple uint

	for multiple = 2 * prime; multiple < limit; multiple += prime {
		list[multiple] = true
	}
}

func get_next_prime(prime uint, list []bool) uint {
	var i uint
	for i = prime + 1; i < limit; i++ {
		if !list[i] {
			return i
		}
	}
	return limit;
}

func get_unmarked_numbers(list []bool) []uint {
	var unmarked []uint
	var i uint
	for i = 0; i < limit; i++ {
		if !list[i] {
			unmarked = append(unmarked, i)
		}
	}

	return unmarked
}

func initialize(numbers []bool) {
	// skip numbers 0 and 1: they are not primes
	numbers[0] = true
	numbers[1] = true
}
