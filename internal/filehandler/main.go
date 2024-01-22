package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close()

	f2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer f2.Close()

	// both read and write --- owner, group, other
	f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
	PrintFatalError(err)
	defer f3.Close()

	scanner := bufio.NewScanner(f3)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println("Found line:", count, scanner.Text())
	}

	f4, err := os.ReadFile("test4.txt")
	PrintFatalError(err)
	context := string(f4)
	fmt.Println(context)


	writebuffer := bufio.NewWriter(f3)
	for i:= 1; i<=5; i++ {
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	// comment the changes
	writebuffer.Flush()
}

// copy file fname1 to fname2
func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	// copy from the second value to first one
	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	// flush file contents to desc
	err = fNew.Sync()
	PrintFatalError(err)



}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
