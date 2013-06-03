goAlfred
========

This go library is used to create workflows for Alfred 2 easier in the <a href="http://golang.org/">go language from Google</a>. It will automatically create your cache and data directories. It also creates the xml listing for the feedback system. Examples included.

The accessible function calls are:

goAlfred.BundleId()  This will get your Bundle Id for your workflow.
goAlfred.Cache()     This function returns the location of your cache directory.
goAlfred.Home()      This function returns the location of your home directory.
goAlfred.Data()      This function returns the location of your workflow's data directory.
goAlfred.Path()      This function returns the location of your workflow's directory.
goAlfred.Error()     Returns the last error received.
AddResult( uid string, arg string, title string, sub string, icon string, valid string, auto string, rtype string)
                     This function allows you to build up the xml string for returning to Alfred.
goAlfred.GetXML()    This function returns the XML string that needs to be given to Alfred.

There are two examples given: feedback.go and mytest.go. The feedback.go program shows how to take an input and return the proper XML sorting to Alfred. The mytest.go is a very simple program to run in a script to returns the input given. The test.alfredworkflow shows how to use the examples in an Alfred workflow.

If you think of anymore functions to include in the library, let me know or fork this library and ask for a pull request.