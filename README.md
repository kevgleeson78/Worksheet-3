# Worksheet-3 Regular Expressions
This is a repository with an example of Regular expressions with the [Go](https://golang.org/) programming language.

Author: [Kevin Gleeson](https://github.com/kevgleeson78)

Third year student at:[GMIT](http://gmit.ie) Galway

## Cloning, compiling and running the application.

1: Download [git](https://git-scm.com/downloads) to your machine if not already installed.

1.1: Download [go](https://golang.org/dl/) if not already installed.

2: Open git bash and cd to the folder you wish to hold the repository.
Alternatively you can right click on the folder and select git bash here.
This will open the git command prompt in the folder selected.
 
 3: To clone the repository type the following command in the terminal making sure you are in the folder needed for the repository.
```bash
>git clone https://github.com/kevgleeson78/Worksheet-3
```
4: To compile the application cd to the folder and type the following 
```bash
> go build 
```
This will compile and create an executable file from the .go file and give it the name of the folder.

5: To run the application ensure you cd to folder the application is held.
Type the following command
```bash
>./worksheet-3
```
6: This will run the application to the terminal window.

## The purpose of this applictaion

The mian purpose of this application is to demonstrate the use of the regexp package with the [Go](https://golang.org/) programming language.
Firstly a function called elizaResponse is created.
A string is passed into this function eg. 
```GO 
fmt.Println(ElizaResponse("I am happy."))
```
The elizaResponse function then takes the input string and checks if the string has matched any regular expression held within the function.
```Go
func Reflections(capturedString string) string {	
```
## Linear process of each function
1: The regexp.Mustcompile(`(?i)I'?\s*a?m(.*)`) function captures everything after I am, I'm or Im, all case ignored with the flag (?i) at the beginning of the expression.
```Go
//I am, I'm, Im all matched with ignore case
//Everything else after I am is captured for the Reflections function
r1 := regexp.MustCompile(`(?i)I'?\s*a?m(.*)`)
//Match the words "I am, Im, I'm" and capture everything else after for replacement
matched := r1.MatchString(str)

```
2: The captured part of the string can the be accessed with the $1 variable this states that it is the first captured group within the expression.
3: The captured string variable $1 is the passed to the regexp function ReplaceAllString() which takes the original string and replaces it with the capturede string.
```Go
//condition if "I am, Im, I'm" is matched
	if matched {
//Only keep the captured part of the string
//Pass in everything after the captured part of the statement $1 to the function Reflections
reflectString := Reflections(r1.ReplaceAllString(str, "$1"))
```
4: The function ReplaceAllString() is in turn passed to the Reflections function which takes each word in the string and splits them into single words by matching spaces between the words and splitting at this point.
5: Using a loop these words are then matched and replaced from a 2d array eg. "I" becomes "you", "You're" becomes "I'm" etc.
```Go
//adapted from https://stackoverflow.com/questions/10196462/regex-word-boundary-excluding-the-hyphen
	//To prevent "you're"  or any word with a "'" from getting split into three tokens
	boundaries := regexp.MustCompile(`(\b[^\w']|$)`)
	//tokens splits up the string into single words for the reflections array to cahge to nouns.
	tokens := boundaries.Split(capturedString, -1)
	// List ofreflections.
	//It depends on the order of the words in the array on how they will be changed from the reflectString function.
	reflections := [][]string{
		{`\byou're\b`, `I'm`},
		{`\b(?i)I'm\b`, `you're`},
		{`\bare\b`, `am`},
		{`\bam\b`, `are`},
		{`\bwere\b`, `was`},
		{`\bwas\b`, `were`},
		{`\byou\b`, `I`},
		{`\bi\b`, `you`},
		{`\bme\b`, `you`},
		{`\byour\b`, `my`},
		{`\bmy\b`, `your`},
		{`\byou've\b`, `I've`},
		{`\bI've\b`, `you've`},
		{`\bme\b`, `you`},
		{`\byou\b`, `me`},
	}
	//Adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
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
	//one continuous string.
	return strings.Join(tokens, " ")
}
```
6: The result is stored in the reflectString Variable.
```Go
reflectString := Reflections(r1.ReplaceAllString(str, "$1"))
```
7: Finally the response is returned with a new message attached to the reflectSting variable.
```Go
response := "Hello " + reflectString + " How are you today?"
		//return the new string response
		return response
```

