package main

//
// Program:              feedback.go
//
// Description:         Create a proper feedback to the alfred system. This is an example of usage.
//

//
// Import the libraries we use for this program.
//
import (
	"./src/goAlfred"
	"fmt"
	"os"
)

//
// Function:          main
//
// Description:       This is the main function that is called whenever the program is
//                           executed. 
//                    
func main() {
	if(len(os.Args) > 1) {
		switch (os.Args[1]) {
			case "1": 
				goAlfred.AddResult("testUID1", "test argument1", "This is my title1", "test substring1", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID2", "test argument2", "This is my title2", "test substring2", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID3", "test argument3", "This is my title3", "test substring3", "icon.png", "yes", "", "")
			case "2":	
				goAlfred.AddResult("testUID2", "test argument2", "This is my title2", "test substring2", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID1", "test argument1", "This is my title1", "test substring1", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID3", "test argument3", "This is my title3", "test substring3", "icon.png", "yes", "", "")
			case "3":
				goAlfred.AddResult("testUID3", "test argument3", "This is my title3", "test substring3", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID1", "test argument1", "This is my title1", "test substring1", "icon.png", "yes", "", "")
				goAlfred.AddResult("testUID2", "test argument2", "This is my title2", "test substring2", "icon.png", "yes", "", "")

		}
	} else {
		goAlfred.AddResult("testUID3", "test argument3", "This is my title3", "test substring3", "icon.png", "yes", "", "")
		goAlfred.AddResult("testUID", "test argument", "This is my title", "test substring", "icon.png", "yes", "", "")
		goAlfred.AddResult("testUID2", "test argument2", "This is my title2", "test substring2", "icon.png", "yes", "", "")		
	}

	//
	// Print out the created XML. 
	//
	fmt.Print(goAlfred.ToXML())
}