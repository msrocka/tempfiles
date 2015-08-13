package main

import (
	"fmt"
	"os"
)

type stat struct {
	size      int64
	fileCount int
	dirCount  int
	failures  int
}

func stats(path string, content []os.FileInfo) {
	fmt.Println("\nStatistics of", path)
	days := dayArg()
	total := stat{}
	old := stat{}
	for _, info := range content {
		s := size(path, info)
		files, dirs := count(path, info)
		older := isOlder(path, info, days)
		total.size += s
		total.fileCount += files
		total.dirCount += dirs
		if older {
			old.size += s
			old.fileCount += files
			old.dirCount += dirs
		}
	}
	total.size /= 1024 * 1024
	old.size /= 1024 * 1024
	printStats(&total, &old, days)
}

func printStats(total *stat, old *stat, days int) {
	fmt.Println("")
	fmt.Printf("  total size:            %d MB\n", total.size)
	fmt.Printf("  total directory count: %d\n", total.dirCount)
	fmt.Printf("  total file count:      %d\n", total.fileCount)
	fmt.Println("")
	fmt.Printf("Of which can be deleted (older %d days):\n\n", days)
	fmt.Printf("  size:                  %d MB\n", old.size)
	fmt.Printf("  directory count:       %d\n", old.dirCount)
	fmt.Printf("  file count:            %d\n", old.fileCount)
	fmt.Println("")
}
