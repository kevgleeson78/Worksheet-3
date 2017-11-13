//worksheet-3
//Regex
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet-3
//date 10/11/2017

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
func ElizaResponse(str string) (string, string) {
	human := str

	random := []string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	matched, _ := regexp.MatchString(`(?i)\bfather\b`, human)
	if matched {

		return human, "Why don’t you tell me more about your father?"
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(random))
	return human, random[randIndex]

}
func main() {
	fmt.Println(ElizaResponse("People say I look like both my mother and father.\n"))
	fmt.Println()
	fmt.Println(ElizaResponse("Father was a teacher.\n"))
	fmt.Println()
	fmt.Println(ElizaResponse("I was my father’s favourite.\n"))
	fmt.Println()
	fmt.Println(ElizaResponse("I’m looking forward to the weekend.\n"))
	fmt.Println()
	fmt.Println(ElizaResponse("My grandfather was French!\n"))
}
