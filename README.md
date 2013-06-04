goAlfred
========

This go library is used to create workflows for Alfred 2 easier in the <a href="http://golang.org/">go language from Google</a>. It will automatically create your cache and data directories. It also creates the xml listing for the feedback system. Examples included.

You install the library with:

go get github.com/raguay/goAlfred     (You might have to use sudo)

Any program that you want to use the library, just place this line in it:

import "github.com/raguay/goAlfred"

The accessible function calls are:
<table>
<tr><td>goAlfred.BundleId()</td><td>This will get your Bundle Id for your workflow.</td></tr>

<tr><td>goAlfred.Cache()</td><td>This function returns the location of your cache directory.</td></tr>

<tr><td>goAlfred.Home()</td><td>This function returns the location of your home directory.</td></tr>

<tr><td>goAlfred.Data()</td><td>This function returns the location of your workflow's data directory.</td></tr>

<tr><td>goAlfred.Path()</td><td>This function returns the location of your workflow's directory.</td></tr>

<tr><td>goAlfred.Error()</td><td>Returns the last error received.</td></tr>
<tr><td>AddResult( uid string, arg string, title string, sub string, icon string, valid string, auto string, rtype string)</td><td>This function allows you to build up the xml string for returning to Alfred.</td></tr>

<tr><td>goAlfred.GetXML()</td><td>This function returns the XML string that needs to be given to Alfred.</td></tr>
</table>

There are two examples given: feedback.go and mytest.go. The feedback.go program shows how to take an input and return the proper XML sorting to Alfred. The mytest.go is a very simple program to run in a script to returns the input given. The test.alfredworkflow shows how to use the examples in an Alfred workflow. These are all in the examples directory.

If you think of anymore functions to include in the library, let me know or fork this library and ask for a pull request.

