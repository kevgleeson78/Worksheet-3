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

#The purpose of this applictaion

The mian purpose of this application is to demonstrate the use of the regexp package with the [Go](https://golang.org/) programming language.
Firstly a function called elizaResponse is created.
A string is passed into this function eg. 
```GO 
fmt.Println(ElizaResponse("I am happy."))
```
The elizaResponse function then takes the input string and checks if the string has matched any regular expression held within the function.
```Go
//I am, I'm, Im all matched with ignore case
r1 := regexp.MustCompile(`(?i)I'?\s*a?m(.*)`)

	//Match the words "I am, Im, I'm" and capture everything else after for replacement
	matched := r1.MatchString(str)

	//condition if "I am, Im, I'm" is matched
	if matched {

		//Only keep the captured part of the string
		//Pass in everything after the captured part of the statement $1 to the function Reflections
		reflectString := Reflections(r1.ReplaceAllString(str, "$1"))
		//Concat the reflected sring to the opening response
		response := "How do you know you are " + reflectString + "?"
		//return the new string response
		return response
    }
    ```

