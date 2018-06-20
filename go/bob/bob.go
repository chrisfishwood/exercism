// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

/*
Bob answers 'Sure.' if you ask him a question.

He answers 'Whoa, chill out!' if you yell at him.

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying
anything.

He answers 'Whatever.' to anything else.
*/

import (
	"fmt"
	"regexp"
	"strings"
)

//} else if strings.HasSuffix(remark, "!") || AllCaps(remark) {

// Hey should have a comment documenting it.
func Hey(remark string) string {
	fmt.Println(remark)
	remark = strings.TrimSpace(remark)
	fmt.Println(remark)
	var response string

	var sayingSomething, sayingWords, yelling, question = IsSayingSomething(remark), IsSayingWords(remark), IsYelling(remark), IsQuestion(remark)
	fmt.Println(sayingSomething, yelling, question)

	if yelling && question && sayingWords {
		response = "Calm down, I know what I'm doing!"
		fmt.Println(response)
	} else if question {
		response = "Sure."
		fmt.Println(response)
	} else if !sayingSomething {
		response = "Fine. Be that way!"
		fmt.Println(response)
	} else if yelling && sayingWords {
		response = "Whoa, chill out!"
		fmt.Println(response)
	} else {
		response = "Whatever."
		fmt.Println(response)
	}

	return response
}

// IsSayingSomething - checks if 'remark' is in all caps
func IsSayingSomething(remark string) bool {
	match, _ := regexp.MatchString("\\w", remark)
	return match
}

// IsSayingWords - checks if 'remark' is in all caps
func IsSayingWords(remark string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]", remark)
	return match
}

// IsYelling - checks if 'remark' is in all caps
func IsYelling(remark string) bool {
	return strings.ToUpper(remark) == remark
	// //match, _ := regexp.MatchString("^[A-Z0-9\\s,]+[!\\.\\?]?$", remark)
	// allCaps := strings.ToUpper(remark) == remark
	// endsWithExcl := strings.HasSuffix(remark, "!")
	// isWords := isWords(remark)
	// //areWords := strings.HasSuffix(remark, "!")
	// fmt.Print("capsMatch:")
	// fmt.Println(allCaps)
	// fmt.Print("endsWith !")
	// fmt.Println(endsWithExcl)
	// fmt.Print("iswords:")
	// fmt.Println(isWords)

	// //match, _ := regexp.MatchString("^([A-Z0-9\\s,]+[!\\.\\?]?$)", remark)
	// fmt.Println("======================")
	// fmt.Println(allCaps && endsWithExcl, isWords)
	// //return match || match && strings.HasSuffix(remark, "!")
	// return (allCaps && endsWithExcl && isWords)
}

// IsWords - is it a question
func isWords(remark string) bool {
	match, _ := regexp.MatchString("^([\\w]+$)", remark)
	return match
}

// IsQuestion - is it a question
func IsQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
