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
*https://regex101.com/r/vH0iN5/2
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

	//pronuon3 := "you"
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
	matchedIAM, _ := regexp.MatchString(`((?i)\bI'?\s*a?m\b)`, str)
	//condition if "I am" is matched
	if matchedIAM {
		r2 := regexp.MustCompile(`(\byou\b)`)

		respronoun := r2.ReplaceAllString(str, "I")
		r2 = regexp.MustCompile(`((?i)\byou'?\s*r?e\b)`)
		respronoun1 := r2.ReplaceAllString(respronoun, "I'm")
		r2 = regexp.MustCompile(`(\byour\b)`)
		respronoun2 := r2.ReplaceAllString(respronoun1, "my")
		//Capture "I am" and ignore case for replacement
		//Also Capture "Im" and "I'm" with ignore case covers different combinations
		r1 := regexp.MustCompile(`((?i)\bI'?\s*a?m\b)`)
		/*create variable to add the response global string
		*to the input string minus the captured expression.
		 */
		firstString := r1.ReplaceAllString(respronoun2, response)
		//Isolate any full stop at the end of each input sentnce
		r1 = regexp.MustCompile(`(\.)`)
		//Second variable to append a "?" and replace any "."
		secondString := r1.ReplaceAllString(firstString, "?")
		//Return the final version of the string
		r1 = regexp.MustCompile(`(me)`)
		thirdString := r1.ReplaceAllString(secondString, "you")
		r1 = regexp.MustCompile(`(I're)`)
		fourthString := r1.ReplaceAllString(thirdString, "I'm")
		return fourthString
	}
	//Match the words "I am" and capture for replacement
	matchedHello, _ := regexp.MatchString(`((?i)\bHello\b)`, str)
	//condition if "I am" is matched
	if matchedHello {
		return "this is working now for the logic :)"
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
	fmt.Println()
	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()
	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()
	fmt.Println("I'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()
	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()
	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()
	fmt.Println("I'm not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I'm not sure that you understand the effect that your questions are having on me."))
	fmt.Println()
	fmt.Println("Im supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("Im supposed to just take what you're saying at face value?"))
	fmt.Println()
	fmt.Println("Hello my name is Kevin")
	fmt.Println(ElizaResponse("Hello my name is Kevin"))
	fmt.Println()
	fmt.Println("How are you feeling")
	fmt.Println(ElizaResponse("How are you feeling"))
	fmt.Println()
	fmt.Println("Im have a hole in my shoe")
	fmt.Println(ElizaResponse("Im supposed"))
	fmt.Println()
}
