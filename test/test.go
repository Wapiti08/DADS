package main

import (
	"fmt"
	"rand"
	"time"
)

func findSC(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)

	if v, ok := scMapping[name]; !ok {
		// return -1, errors.New("Crew member not found")
		return -1, fmt.Errorf("Crew member not found")

	} else {
		return v, nil
	}
}

func main() {
	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(s)
	s1 := make([]byte, 5, 5)
	fmt.Println(s1)
	fmt.Println(len(s1))
	s2 := []int{1, 2, 3}
	s2 = append(s2, 4, 5, 6)
	s3 := []int{4, 5, 6}
	// convert slice s3 into elements
	s3 = append(s3, s3...)
	fmt.Println(s3)
}
