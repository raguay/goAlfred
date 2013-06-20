package goAlfred

//
// Package:          goAlfred
//
// Description:      This package is for helper function to create workflows for Alfred2 using
//                           go. 
//

//
// Import the libraries we use for this library.
//
import (
	"log"
	"bytes"
	"os"
	"encoding/xml"
	"io/ioutil"
	"github.com/mkrautz/plist"
	"strings"
	"regexp"
)

//
// Class Variables
//
// Name                               Description
//
// cache                                path to the directory that contains the cache for the workflow
// data                                  path to the directory that contains the data for the workflow
// bundleId                           The ID for the bundle that represents the workflow
// path                                  path to the workflow's directory
// home                                path to the user's home directory
// results                              the accumulated results. This will be converted to the XML list for 
//                                          feedback into Alfred
// err                                     The value of the last error found
//

type AlfredResult struct {
	XMLName   xml.Name `xml:"item"`	
	Uid              string       `xml:"uidid,attr"`
	Arg              string       `xml:"arg"`
	Title             string       `xml:"title"`
	Sub              string       `xml:"sub"`
	Icon             string       `xml:"icon"`
	Valid            string          `xml:"valid,attr"`
	Auto             string       `xml:"auto,attr"`
	Rtype            string       `xml:"type,attr,omitempty"`
}

var(
	cache string
	data string
	bundleId string
	path string
	home string
	err error
	maxResults int
	currentResult int
	results []AlfredResult
)

//
// Library Function:
//
//				init 			This function is called upon library use to initialize 
//							any variables used for the library before anyone
// 							can make a call to a library function. 
//
func init() {
	//
	// Set the path and home variables from the environment.
	//
	path, err = os.Getwd()
	home =  os.Getenv("HOME")
	if _, err = os.Stat("info.plist" ); err == nil {
		//
		// The file exists. Read it for the bundleid and set the bundleId variable. 
		//
		bundleId = GetBundleId()
	} else {
		//
		// Give an error message and set error to it. Then return. 
		//
		log.Println("There is no plist!")
		return
	}

	//
	// Create the directory structure for the cache and data directories. 
	//
	cache = home + "/Library/Caches/com.runningwithcrayons.Alfred-2/Workflow Data/" + bundleId;
	data  = home + "/Library/Application Support/Alfred 2/Workflow Data/" + bundleId;

	//
	// See if the cache directory exists. 
	//
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		//
		// The cache directory does not exist. Create it. 
		//
		err = os.MkdirAll(cache, 0777|os.ModeDir)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = os.Chmod(cache, 0777|os.ModeDir)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	//
	// See if the data directory exists. 
	//
	if _, err := os.Stat(data); os.IsNotExist(err) {
		//
		// The data directory does not exist. Create it. 
		//
		err = os.MkdirAll(data, 0777|os.ModeDir)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = os.Chmod(data, 0777|os.ModeDir)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	//
	// Create the result array. 
	//
	results = make([]AlfredResult, 10)
	results[0].Title = "No matches found..."
	results[0].Uid = "default"
	results[0].Valid = "no"
	maxResults = 10
	currentResult = 0
}

//
// Function:           GetBundleId 
//
// Description:       This function will read the workflows info.plist and return
//                            the bundleid
//
func GetBundleId() string {
	//
	// My version before the plist reader worked. 
	//
	fileloc := path + "/info.plist"
	//myout, err := exec.Command("/usr/bin/defaults", "read", fileloc , "bundleid").Output()
	//if(err != nil) {
	//	log.Fatalf("Error with command: %v" ,err)
	//}
	//return(string(myout))

	buf, err := ioutil.ReadFile(fileloc)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var workflow map[string]interface{}
	err = plist.Unmarshal(buf, &workflow)
	if err != nil {
		log.Fatalf("%v", err)
	}

	v, ok := workflow["bundleid"]
	if !ok {
		log.Fatalf("expected bundleid key, but wasn't found")
	}

	//
	// Return the bundle ID.
	//
	return(v.(string))
}

//
// Function:           BundleId 
//
// Description:       This function returns the bundleid for the workflow. 
//

func BundleId() string {
	return(  bundleId )
}

//
// Function:           Cache 
//
// Description:       This function returns the cache directory for the workflow. 
//
func Cache() string {
	return(  cache )  
}

//
// Function:           Data 
//
// Description:       This function returns the data directory for the workflow. 
//
func Data() string {
	return(data)
}

//
// Function:           Path 
//
// Description:       This function returns the path to the workflow. 
//
func Path() string {
	return(path)
}

//
// Function:           Home 
//
// Description:       This function returns the Home directory for the user. 
//
func Home() string {
	return(home)
}

//
// Function:           Error 
//
// Description:       This routine will return the error string. 
//
func Error() error {
	return(err)
}

//
// Function:           ToXML 
//
// Description:       This function takes the result array and makes it into an
//                            XML string for passing back to Alfred.  Possible help: 
//                            http://golang.org/pkg/encoding/xml/#example_MarshalIndent
//
// Inputs:
//                             arg          A string to base the ordering. 
//
func ToXML() string {
	//
	// Initialize the output string and create a string writer. 
	//
	newxml := "<items>"
	buf := bytes.NewBufferString(newxml)

	//
	// Create the xml encoder. 
	//
	enc := xml.NewEncoder(buf)

	//
	// Encode it. If there is an error, print it to the log. 
	//
	//if err := enc.Encode(results); err != nil {
	//	log.Fatalf("ToXML Error: %v\n", err)
	//}
	for i:=0;i<maxResults;i++ {
		if results[i].Uid != "" {
			if err := enc.Encode(results[i]); err != nil {
				log.Fatalf("ToXML Error: %v\n", err)	
			}
		}
	}

	//
	// Convert the buffer to a string and add the closing tag. 
	//
	newxml = buf.String() + "</items>\n"

	//
	// Return the XML string. 
	//
	return(newxml)
}

//
// Function:           AddResult 
//
// Description:       Helper function that just makes it easier to pass values into a function
//                           and create an array result to be passed back to Alfred. 
//
// Inputs:
// 		uid 		the uid of the result, should be unique
// 		arg 		the argument that will be passed on
// 		title 		The title of the result item
// 		sub 		The subtitle text for the result item
// 		icon 		the icon to use for the result item
// 		valid 		sets whether the result item can be actioned
// 		auto 		the autocomplete value for the result item
//              rtype           I have no idea what this one is used for. HELP!
//
func AddResult( uid string, arg string, title string, sub string, icon string, valid string, auto string, rtype string) {
	//
	// Add in the new result array if not full. 
	//
	if (currentResult < maxResults) {
		results[currentResult] .Uid = uid
		results[currentResult] .Arg = arg
		results[currentResult] .Title = title
		results[currentResult] .Sub = sub
		results[currentResult].Icon = icon
		results[currentResult] .Valid = valid
		results[currentResult] .Auto = auto
		results[currentResult].Rtype = rtype
		currentResult++
	}
}

//
// Function:           AddResultsSimilar 
//
// Description:       This function will only add the results that are similar to the input given. 
//                            This is used to select input selectively from what the user types in. 
//
// Inputs:
//               instring      the string to test against the titles to allow that record or not
// 		uid 		the uid of the result, should be unique
// 		arg 		the argument that will be passed on
// 		title 		The title of the result item
// 		sub 		The subtitle text for the result item
// 		icon 		the icon to use for the result item
// 		valid 		sets whether the result item can be actioned
// 		auto 		the autocomplete value for the result item
//              rtype           I have no idea what this one is used for. HELP!
//
func AddResultsSimilar( instring string, uid string, arg string, title string, sub string, icon string, valid string, auto string, rtype string) {
	//
	// Create the test pattern. 
	//
	instring = strings.ToLower(instring) + ".*"

	// 
	// Compare the match string to the title for the Alfred output. 
	//
         mt, _ := regexp.MatchString(instring, strings.ToLower(title))
         if(mt) {
         	//
         	// A match, add it to the results. 
         	//
         	AddResult( uid, arg, title, sub, icon, valid,  auto, rtype) 
         }
}

//
// Function:           SetDefaultString
//
// Description:       This function sets a different default title
//
// Inputs:
// 		title 	the title to use
//
func SetDefaultString(title string) {
	if(currentResult == 0) {
		//
		// Add only if no results have been added.
		//
		results[0].Title = title
	}
}