package rsa

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RSASuite struct{}

var _ = Suite(&RSASuite{})

func (r *RSASuite) Test_isPrime_ReturnsTrueWhenASimplePrimeNumberIsGiven(c *C) {
	number := 5
	result := isPrime(number)
	c.Assert(result, Equals, true)
}

func (r *RSASuite) Test_isPrime_ReturnsFalseWhenNoPrimeNumberIsGiven(c *C) {
	number := 10
	result := isPrime(number)
	c.Assert(result, Equals, false)
}
