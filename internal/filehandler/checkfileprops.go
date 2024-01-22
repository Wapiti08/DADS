package main

import (
	"fmt"
	"os"
)

// func main() {
// 	GenerateFileStatusReport("testfile.txt")
// }

func GenerateFileStatusReport(fname string) {
	// Stat return file info
	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("What is the file name?", filestats.Name())
	fmt.Println("Am I a directory?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("What's the file size?", filestats.Size())
	fmt.Println("What was the last time the file modified?", filestats.ModTime())

}
