package luhn

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Valid -
func Valid(word string) bool {
	//strip spaces
	//if length <= 1 return false
	//non-digits are invalid
	word = strings.TrimSpace(word)
	fmt.Print("trimmed word:")
	fmt.Println(word)
	if len(word) <= 1 {
		return false
	}
	//059
	isOdd := len(word)%2 == 1
	fmt.Print("before:")
	fmt.Println(word)
	word = trimAllSpace(word)
	fmt.Print("after:")
	fmt.Println(word)
	var newString bytes.Buffer

	for i := 0; i < len(word); i++ {
		stringDigit := string(word[i])
		if (isOdd && i%2 == 0) || stringDigit == " " {
			fmt.Printf("%s is odd, skipping\n", stringDigit)
			newString.WriteString(stringDigit)
			continue
		}
		intDouble, _ := strconv.Atoi(stringDigit)
		fmt.Printf("int pre doubling: %d \n", intDouble)
		intDouble = intDouble * 2
		fmt.Printf("int post doubling: %d \n", intDouble)
		if intDouble > 9 {
			intDouble = intDouble - 9
		}
		fmt.Printf("int post >9 check: %d \n", intDouble)
		newString.WriteString(strconv.FormatInt(int64(intDouble), 10))
	}

	fmt.Printf("New String: %s\n", newString.String())

	reversedWord := newString.String()
	sum := 0
	for i := 0; i < len(reversedWord); i++ {
		sumInt, _ := strconv.Atoi(string(reversedWord[i]))
		sum = sum + sumInt
	}
	/* 	revValue := ReverseWord(word)
	   	fmt.Print("reversed word:")
	   	fmt.Println(revValue)

	   	for i := 0; i < len(revValue); i++ { // i := 1 to skip the first char
	   		digit, _ := strconv.ParseInt(string(revValue[i]), 10, 32)

	   		fmt.Print("single digit:")
	   		fmt.Println(revValue)
	   		if isOdd(digit) {
	   			//double it

	   		}
	   	}*/
	return sum%10 == 0
}

func trimAllSpace(word string) string {
	var noSpaces bytes.Buffer

	for i := 0; i < len(word); i++ {
		char := string(string(word[i]))
		fmt.Print("in trim spaces")
		fmt.Println(char)
		if char != " " {
			noSpaces.WriteString(char)
		}
	}

	return noSpaces.String()
}

func isOdd(num int64) bool {
	return num%2 == 0
}

//ReverseWord - Takes a string and returns it reverse order
func reverseWord(word string) string {

	fmt.Println("reverse")
	var reverse bytes.Buffer

	for i := len(word) - 1; i != 0; i-- {
		fmt.Print("in reverse")
		fmt.Println(string(word[i]))
		reverse.WriteByte(word[i])
	}

	return reverse.String()
}
