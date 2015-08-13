package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTmpDir() string {
	candidates := [...]string{"TEMP", "TMP", "TEMPORARY"}
	for _, candidate := range candidates {
		env := os.Getenv(candidate)
		if env != "" {
			return env
		}
	}
	return "Not found!"
}

// Visits the given file in the given path. If the given file is a
// directory it visits also the directory content recursively.
func visit(path string, info os.FileInfo, fn func(os.FileInfo)) error {
	fn(info)
	if !info.IsDir() {
		return nil
	}
	fullPath := path + string(os.PathSeparator) + info.Name()
	list, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return err
	}
	for _, next := range list {
		err = visit(fullPath, next, fn)
		if err != nil {
			return err
		}
	}
	return nil
}

// Returns the age of the given file in days.
func age(info os.FileInfo) int {
	now := time.Now()
	diff := now.Sub(info.ModTime())
	hours := diff.Hours()
	return int(hours) / 24
}

// Returns the information about the file that should be displayed
// in the console.
func infoString(info os.FileInfo) string {
	name := info.Name()
	age := strconv.Itoa(age(info)) + " days"
	if info.IsDir() {
		return name + " (dir, " + age + ")"
	} else {
		size := strconv.Itoa(int(info.Size())) + " bytes"
		return name + " (" + size + ", " + age + ")"
	}
}

// Returns the size of the given file. If the given file is a directory
// it recursively calculates the size from the directory content.
func size(path string, info os.FileInfo) int64 {
	if !info.IsDir() {
		return info.Size()
	}
	var total int64 = 0
	visit(path, info, func(next os.FileInfo) {
		if !next.IsDir() {
			total += next.Size()
		}
	})
	return total
}

func count(path string, info os.FileInfo) (files, dirs int) {
	if !info.IsDir() {
		return 1, 0
	}
	files = 0
	dirs = 0
	visit(path, info, func(next os.FileInfo) {
		if next.IsDir() {
			dirs++
		} else {
			files++
		}
	})
	return files, dirs
}

func isOlder(path string, info os.FileInfo, days int) bool {
	if !info.IsDir() {
		return age(info) > days
	}
	deletable := true
	visit(path, info, func(next os.FileInfo) {
		if age(next) <= days {
			deletable = false
		}
	})
	return deletable
}

func dayArg() int {
	if len(os.Args) < 3 {
		return DEFAULT_DAYS
	}
	s := os.Args[2]
	if !strings.HasPrefix(s, "days=") {
		return DEFAULT_DAYS
	}
	s = strings.TrimPrefix(s, "days=")
	d, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("could not read days %s; took default = %d\n", s, DEFAULT_DAYS)
		return DEFAULT_DAYS
	}
	if d < 0 {
		fmt.Printf("%d < 0; took default = %d\n days", d, DEFAULT_DAYS)
		return DEFAULT_DAYS
	}
	return d
}
