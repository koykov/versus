package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	reChar = regexp.MustCompile(`<char cp="([0-9A-Z]+)"`)
)

func main() {
	file, err := os.Open("./testdata/ucd.all.flat.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	scr := bufio.NewScanner(file)
	for scr.Scan() {
		line := scr.Text()
		if len(line) == 0 {
			continue
		}
		if m := reChar.FindStringSubmatch(line); len(m) > 0 {
			// fmt.Println(m[1])
		} else {
			fmt.Println(line)
		}
	}

	if err := scr.Err(); err != nil {
		log.Fatal(err)
	}
}
