package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const DEFAULT_DAYS = 5

func printHelp() {
	h := `
    Usage of tempfiles:

    clean [days=DAYS]: deletes the temporary files and folders that are older 
                       than DAYS (default = 5)

    dir:   prints the temporary file folder

    env:   prints all environment variables as KEY=VALUE pairs
	
	list:  lists the content of the temporary file folder

    stats [days=DAYS]: calculates and prints statistics about the temporary 
                       file folder in total and regarding the files that can
                       be deleted

    help:  prints this help`
	fmt.Println(h)
}

func printEnv() {
	env := os.Environ()
	for _, line := range env {
		fmt.Println(line)
	}
}

type contentHandler func(path string, content []os.FileInfo)

func run(fn contentHandler) {
	dir := getTmpDir()
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("failed to read files from", dir, err)
		return
	}
	fn(dir, list)
}

func list(path string, content []os.FileInfo) {
	fmt.Println("Content of", path)
	for _, info := range content {
		if info.IsDir() {
			visit(path, info, func(next os.FileInfo) {
				fmt.Println(info.Name() + " .. " + infoString(next))
			})
		} else {
			fmt.Println(infoString(info))
		}
	}
}

func main() {
	argLen := len(os.Args)
	if argLen < 2 {
		printHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "env":
		printEnv()
	case "help":
		printHelp()
	case "dir":
		fmt.Println(getTmpDir())
	case "list":
		run(list)
	case "clean":
		run(clean)
	case "stats":
		run(stats)
	default:
		fmt.Println("unknown commad", command)
		printHelp()
	}

}
