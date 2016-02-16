/* Counts the number of primes less than or equal to x (usually denoted by
 * pi(x)), x > 0, using the crude approximation given by the prime number
 * theorem (x/lnx). */

package main

var x float64 = 100
var lnx, term float64
var niter int = 100

func main() {
	// Approximate lnx using the first 100 terms of the Mercator series by
	// iterating backwards over the terms.
	for i := 0; i < niter; i++ {
		term = 1.

		// Calculate (x - 1)^i.
		for j := 0 ; j < i ; j++ {
			term = term * (x - 1)
		}

		// (x - 1)^i / i.
		term = term / i;

		// Substract if its an even numbered term in the sequence, add otherwise
		if i % 2 == 0 {
			lnx = lnx - term
		} else {
			lnx = lnx + term
		}
	}

	println(x / lnx)
}
