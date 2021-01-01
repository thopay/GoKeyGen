package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var qty int
	var fileName, format string
	keys := []string{}
	nums := []rune("0123456789")
	upperCase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerCase := []rune("abcdefghijklmnopqrstuvwxyz")
	any := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	symbols := []rune("}\\/!$%&+-,=(){")
	test := [5]rune{'?', '&', '.', '*', '%'}
	fmt.Print(`Formatting:
Numbers    :  ?
Uppercase  :  &
Lowercase  :  .
a-z|A-Z|0-9:  *
Any Symbol :  %
Number of Keys: `)
	fmt.Scanln(&qty)
	fmt.Print("Output File Name: ")
	fmt.Scanln(&fileName)
	if !strings.Contains(fileName, ".txt") {
		fileName = fileName + ".txt"
	}
	fmt.Print("Key Format: ")
	fmt.Scanln(&format)
	start := time.Now()
	for len(keys) < qty {
		a := ""
		for _, b := range format {
			if b == test[0] {
				a += string(nums[rand.Intn(len(nums))])
			} else if b == test[1] {
				a += string(upperCase[rand.Intn(len(upperCase))])
			} else if b == test[2] {
				a += string(lowerCase[rand.Intn(len(lowerCase))])
			} else if b == test[3] {
				a += string(any[rand.Intn(len(any))])
			} else if b == test[4] {
				a += string(symbols[rand.Intn(len(symbols))])
			} else {
				a += string(b)
			}
		}
		keys = append(keys, a)
	}
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	for _, b := range keys {
		_, err2 := f.WriteString(b + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("Done in", time.Since(start))
}
