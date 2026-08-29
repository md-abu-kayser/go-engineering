// Lesson 20: Logical Operators
//
// Goal: Combine COMPARISON expressions with &&/||/! into more complex
// conditions, and apply De Morgan's laws to simplify negated logic.
package main

import "fmt"

func canVote(age int, isCitizen bool) bool {
	return age >= 18 && isCitizen
}

func main() {
	fmt.Println("=== Logical Operators ===")
	fmt.Println("----------------------------------")

	fmt.Printf("canVote(20, true)  : %t\n", canVote(20, true))
	fmt.Printf("canVote(20, false) : %t\n", canVote(20, false))
	fmt.Printf("canVote(15, true)  : %t\n", canVote(15, true))

	// Combining THREE conditions — precedence (lesson 17) means && binds
	// tighter than ||, so this reads as: cond1 || (cond2 && cond3).
	age, hasLicense, hasPermit := 16, false, true
	canDrive := age >= 18 || (hasLicense && hasPermit)
	fmt.Printf("\ncanDrive (16, no license, has permit) : %t\n", canDrive)

	// De Morgan's laws: negating a combined condition lets you "push"
	// the ! inward, flipping && to || (or vice versa) and negating
	// each individual condition.
	fmt.Println("\n--- De Morgan's laws ---")
	a, b := true, false

	notBoth := !(a && b)
	deMorgan1 := !a || !b
	fmt.Printf("!(a && b)      = %t\n", notBoth)
	fmt.Printf("!a || !b       = %t (equivalent, by De Morgan's law)\n", deMorgan1)

	notEither := !(a || b)
	deMorgan2 := !a && !b
	fmt.Printf("!(a || b)      = %t\n", notEither)
	fmt.Printf("!a && !b       = %t (equivalent, by De Morgan's law)\n", deMorgan2)
}
