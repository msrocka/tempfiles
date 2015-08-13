#tempfiles for Windows
This is a small command line utility for deleting temporary files in Windows. In
Windows temporary files are not deleted by default and you have to do this 
manually. This tool can help you to do this faster.

##Usage
Just take the tempfiles executable and execute it with one of the following 
command (e.g. `tempfiles list`):

    clean [days=DAYS]: deletes the temporary files and folders that are older 
                       than DAYS (default = 5)

    dir:   prints the temporary file folder

    env:   prints all environment variables as KEY=VALUE pairs
	
	list:  lists the content of the temporary file folder

    stats [days=DAYS]: calculates and prints statistics about the temporary 
                       file folder in total and regarding the files that can
                       be deleted

    help:  prints the help

You may want to put the tempfiles executable into a directory which is in your 
system path so that you can execute it from 

##Building from source
This is a plain [Go](https://golang.org/) project. You can just get the source 
via

    go get github.com/msrocka/tempfiles

and build the executable

    go install github.com/msrocka/tempfiles
