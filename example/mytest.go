package main

//
// Program:              program.go
//
// Description:         This is a basic program shell for Alfred.  It is an example for creating your own. 
//

//
// Import the libraries we use for this program.
//
import (
	"github.com/raguay/goAlfred"
	"os"
	"fmt"
)

func main() {
	if(len(os.Args) >1) {
		fmt.Print("I was given: " + os.Args[1])
	} else {
		fmt.Print("Nothing was given.")
	}
	fmt.Printf("The workflow directory is: %s\n", goAlfred.Path())
	fmt.Printf("The bundle id is: %s\n", goAlfred.BundleId())
}