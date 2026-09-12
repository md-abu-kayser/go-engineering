// Lesson 14: Chain Of Responsibility
//
// Goal: Make a variation point explicit with a small interface, then choose
// a concrete policy at the composition point.
package main

import "fmt"

const lesson = "Chain Of Responsibility"

type discount interface {
	Apply(int) int
}

type percentageDiscount struct {
	percent int
}

func (d percentageDiscount) Apply(price int) int {
	return price - price*d.percent/100
}

func checkout(policy discount, price int) int {
	return policy.Apply(price)
}

func main() {
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("discounted total: %d\n", checkout(percentageDiscount{percent: 15}, 200))
}