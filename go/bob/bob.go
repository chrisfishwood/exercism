package bob

/* Checks that Bob replies correctly when asked a variety of questions.
If 'remark' is yelling, is a question and is sayingWords then it replies with:
	"Calm down, I know what I'm doing!"
If 'remark' is a question then it replies with:
	"Sure."
If 'remark' is a silence (no characters) then it replies with:
	"Fine. Be that way!"
If 'remark' is a yelling actual words then it replies with:
	"Whoa, chill out!"
Everything else replies with:
	"Whatever."
*/

import (
	"regexp"
	"strings"
)

// Hey - executes the Bob-rules for replies as defined above.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var response string
	var silence, sayingWords, yelling, question = isEmpty(remark), isSayingWords(remark), isYelling(remark), isQuestion(remark)

	if yelling && question && sayingWords {
		response = "Calm down, I know what I'm doing!"
	} else if question {
		response = "Sure."
	} else if silence {
		response = "Fine. Be that way!"
	} else if yelling && sayingWords {
		response = "Whoa, chill out!"
	} else {
		response = "Whatever."
	}

	return response
}

// isEmpty - checks if 'remark' contains any characters
func isEmpty(remark string) bool {
	return remark == ""
}

// isSayingWords - checks if 'remark' are word characters
func isSayingWords(remark string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]", remark)
	return match
}

// isYelling - checks if 'remark' is in all caps
func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark
}

// isWords - contains word charaters
func isWords(remark string) bool {
	match, _ := regexp.MatchString("^([\\w]+$)", remark)
	return match
}

// isQuestion - Checks if remark ends with a '?'
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
