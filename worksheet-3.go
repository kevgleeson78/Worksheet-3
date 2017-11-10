//worksheet-3
//Regex
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet-3
//date 10/11/2017

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Function ElizaResponse that takes in and returns a Random String
*by referencing its array index.
 */
func ElizaResponse() string {

	rand.Seed(time.Now().UnixNano())
	random := []string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}
	return random[rand.Intn(3-0)]
}
func main() {
	fmt.Println(ElizaResponse())
}
