package main

//Lukas Siemers
//Extra Credit Assignment
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
I ask the User for Input and Initialize the var Input before the loop with that Variable, depending on the answer of the user
we will capitialize the input so that we don't have problems with the spelling. if yes is selected then we send a go function call as well as
call the channel with a boolean which is initialized to true and is being used inside the function.
Inside the function if the value from the channel is true we are opening the file and check for errors.
We check that the file is closed at the end of the function and use bufio as well as Strings.Builder to help us structure
the seperation from the special commands and each individual line. After that we go ahead and initialize fortune with the split string and
create a "random" variable at the length of the Quotes. We then use the Random number as an index to find a qoute and print it off.
At the end of the function we have another channel which allows us to determine when we are done using the function and returns us to the main
where we ask the user for more input or to end the program with no.
*/

func fortune(ch <-chan bool, done chan<- bool) {
	value := <-ch // receive value from channel ch
	for value == true {
		// Open the file
		file, err := os.Open("C:\\Users\\Lukas\\IdeaProjects\\ExtraCredit\\Fortunes.txt") //Modify the TextFile location since I have it locally on my Desktop
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()           // Ensure file is closed at the end of the function
		sc := bufio.NewScanner(file) // Use bufio.Scanner to read the file
		var content strings.Builder
		for sc.Scan() {
			content.WriteString(sc.Text() + "\n") // Add each line to the string builder
		}
		if err := sc.Err(); err != nil { // Check for errors in scanning
			log.Fatal(err)
		}
		fortunes := strings.Split(content.String(), "%%") // Split the entire content by '%%'
		rand.Seed(time.Now().UnixNano())                  // Initialize random seed
		if len(fortunes) > 0 {                            // Select a random fortune
			randomIndex := rand.Intn(len(fortunes))
			fmt.Println(strings.TrimSpace(fortunes[randomIndex])) // Print the randomly selected fortune
		}
		done <- true //Allows me to return to the Main function and ask for another Input
		return
	}
}

func main() {
	var input string
	done := make(chan bool)                       // channel to signal main the task is done
	ch := make(chan bool)                         //Create a channel
	fmt.Print("would you like another fortune? ") //Ask for Initial Input outside the loop
	fmt.Scan(&input)                              //Scan for user Input
	for strings.ToUpper(input) != "NO" {
		if strings.ToUpper(input) == "YES" {
			go fortune(ch, done)
			ch <- true
			<-done
			fmt.Print("would you like another fortune? ")
			fmt.Scan(&input)
		} else {
			fmt.Print("would you like another fortune? ")
			fmt.Scan(&input)
		}
	}
	close(ch)
	fmt.Println("Program has ended")
}
