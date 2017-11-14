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
*
*
 */
package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

/*Function ElizaResponse that takes in and returns a Random String
*by referencing its array index.
 */
func ElizaResponse(str string) string {

	response := "How do you know you are"
	random := []string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	matched, _ := regexp.MatchString(`(?i)\bfather\b`, str)
	if matched {
		return "Why don’t you tell me more about your father?"
	}
	matchedIAM, _ := regexp.MatchString(`I am`, str)
	if matchedIAM {
		r1 := regexp.MustCompile(`((?i)I Am)`)
		firstString := r1.ReplaceAllString(str, response)
		r1 = regexp.MustCompile(`(\.)`)
		secondString := r1.ReplaceAllString(firstString, "?")
		return secondString
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(random))
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
