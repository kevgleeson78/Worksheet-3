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
*https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
*https://stackoverflow.com/questions/10196462/regex-word-boundary-excluding-the-hyphen
*https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
 */
package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//Function ElizaResponse to take in and return a string
func ElizaResponse(str string) string {
	//Adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
	//	replace := "How do you know you are"

	/*Regex MatchString function with isolation of the word "father"
	*with a boundry ignore case regex command.
	 */
	if matched, _ := regexp.MatchString(`(?i)\bfather\b`, str);
	//Condition to replace the original string if it has the word "father"
	matched {
		return "Why don’t you tell me more about your father?"
	}
	r1 := regexp.MustCompile(`(?i)\bI'?\s*a?m\b`)

	//Match the words "I am" and capture for replacement
	matched := r1.MatchString(str)

	//condition if "I am" is matched
	if matched {
		//Only keep the captured part of the string
		//Concat the new opening line at the end of the function
		capturedString := r1.ReplaceAllString(str, "$1")
		//adapted from https://stackoverflow.com/questions/10196462/regex-word-boundary-excluding-the-hyphen
		//To prevent "you're"  or any word with a "'" from getting split into three tokens
		boundaries := regexp.MustCompile(`(\b[^\w']|$)`)
		tokens := boundaries.Split(capturedString, -1)
		fmt.Println(tokens)
		// List the reflections.
		reflections := [][]string{
			{`you're`, `I'm`},
			{`I`, `you`},
			{`your`, `my`},
			{`me`, `you`},
			{`you`, `I`},
			{`my`, `your`},
			{`\.`, `?`},
		}

		// Loop through each token, reflecting it if there's a match.
		for i, token := range tokens {
			for _, reflection := range reflections {
				if matched, _ := regexp.MatchString(reflection[0], token); matched {
					tokens[i] = reflection[1]
					break
				}
			}
		}

		// Put the tokens back together.
		//A space is need for teh regular expression (\b[^\w']|$)
		//as it dosent allow the word you're to be split into three parts.
		//If the space is not put in as the second argument it will return
		//one continious string.
		return strings.Join(tokens, ` `)

	}

	//Get random number from the length of the array of random struct
	//an array of strings for the random response
	response := []string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}
	//Return a random index of the array
	return response[rand.Intn(len(response))]

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
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
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
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
