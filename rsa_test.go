package rsa

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RSASuite struct{}

var _ = Suite(&RSASuite{})

func (r *RSASuite) Test_isPrime_ReturnsTrueWhenASimplePrimeNumberIsGiven(c *C) {
	number1 := 11
	number2 := 3
	number3 := 5
	result1 := isPrime(number1)
	result2 := isPrime(number2)
	result3 := isPrime(number3)
	c.Assert(result1, Equals, true)
	c.Assert(result2, Equals, true)
	c.Assert(result3, Equals, true)
}

func (r *RSASuite) Test_isPrime_ReturnsFalseWhenNoPrimeNumberIsGiven(c *C) {
	number1 := 12
	number2 := 30
	number3 := 25
	result1 := isPrime(number1)
	result2 := isPrime(number2)
	result3 := isPrime(number3)
	c.Assert(result1, Equals, false)
	c.Assert(result2, Equals, false)
	c.Assert(result3, Equals, false)
}
