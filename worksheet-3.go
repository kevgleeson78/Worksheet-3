//worksheet-3
//Regex
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet-3
//date 10/11/2017
//Version:  1.0
/*Sources:
*https://gobyexample.com/regular-expressions
*https://regexr.com/
*https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.3.html
*https://regexone.com/lesson/letters_and_digits?
*https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter1.markdown
*
 */
package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//Function to take in and return a string
func ElizaResponse(str string) string {
	//Global response variable for problem 3
	response := "How do you know you are"
	//an array of strings for the random response
	random := []string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}
	/*Regex MatchString function with isolation of the word father
	*in comparison to the input string to the function ElizaResponse
	 */
	matched, _ := regexp.MatchString(`(?i)\bfather\b`, str)
	//Condition to replace the original string if it has the word "father"
	if matched {
		return "Why don’t you tell me more about your father?"
	}
	//Match the words "I am" and capture for replacement
	matchedIAM, _ := regexp.MatchString(`I am`, str)
	//condition if "I am" is matched
	if matchedIAM {
		//Capture "I am" and ignore case for replacement
		r1 := regexp.MustCompile(`((?i)I Am)`)
		/*create variable to add the response global string
		*to the input string minus the captured expression.
		 */firstString := r1.ReplaceAllString(str, response)
		//Isolate any full stop at the end of each input sentnce
		r1 = regexp.MustCompile(`(\.)`)
		//Second variable to append a "?" and replace any "."
		secondString := r1.ReplaceAllString(firstString, "?")
		//Return the final version of the string
		return secondString
	}
	//Get time to seed a number to ensure a valid random sequence
	rand.Seed(time.Now().UnixNano())
	//Get random number from the length of the array of random struct
	randIndex := rand.Intn(len(random))
	//Return a random index of the array
	return random[randIndex]
}
func main() {
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()
}
