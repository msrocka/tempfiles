package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func clean(path string, content []os.FileInfo) {
	days := dayArg()
	fmt.Printf("\nDelete everything that is older than %d days?[y/n]\n", days)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil || strings.ToLower(strings.TrimSpace(s)) != "y" {
		fmt.Println("Nothing done")
		return
	}
	deleted := stat{}
	for _, info := range content {
		if !isOlder(path, info, days) {
			continue
		}
		size := size(path, info)
		files, dirs := count(path, info)
		fullPath := path + string(os.PathSeparator) + info.Name()
		err = os.RemoveAll(fullPath)
		if err != nil {
			fmt.Println("Failed to delete", fullPath)
			deleted.failures++
		} else {
			fmt.Println("Deleted", fullPath)
			deleted.dirCount += dirs
			deleted.fileCount += files
			deleted.size += size
		}
	}
	deleted.size /= (1024 * 1024)
	printDeletedStat(&deleted)

}

func printDeletedStat(deleted *stat) {
	fmt.Printf("Deleted %d MB, %d files, %d directories\n", deleted.size,
		deleted.fileCount, deleted.dirCount)
	fmt.Printf("%d failures\n", deleted.failures)
}
