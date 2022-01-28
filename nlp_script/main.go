package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
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

	var (
		i64 int64
		buf bytes.Buffer
		ln  []uint
	)

	_, _ = buf.WriteString("// DO NOT EDIT.\n\npackage testdata\n\ntype Tuple struct{\n\tC int32\n\tS []uint\n}\n\nvar Model = []Tuple{\n")

	scr := bufio.NewScanner(file)
	for scr.Scan() {
		line := scr.Text()
		if len(line) == 0 {
			continue
		}
		if m := reChar.FindStringSubmatch(line); len(m) > 0 {
			if i64, err = strconv.ParseInt(m[1], 16, 64); err != nil {
				fmt.Println(err)
				continue
			}
			r := rune(i64)
			ln = ln[:0]
			for i := 0; i < len(scripts); i++ {
				s := &scripts[i]
				if unicode.Is(s.b, r) {
					ln = append(ln, uint(i))
				}
			}
			if len(ln) > 0 {
				_ = buf.WriteByte('{')
				_, _ = buf.WriteString(strconv.Itoa(int(i64)))
				_, _ = buf.WriteString(", []uint{")
				for i := 0; i < len(ln); i++ {
					if i > 0 {
						_ = buf.WriteByte(',')
					}
					_, _ = buf.WriteString(strconv.Itoa(int(ln[i])))
				}
				_, _ = buf.WriteString("}},\n")
			}
		}
	}

	if err := scr.Err(); err != nil {
		log.Fatal(err)
	}

	_, _ = buf.WriteString("}\n")

	source := buf.Bytes()
	var fmtSource []byte
	if fmtSource, err = format.Source(source); err != nil {
		return
	}

	if err = ioutil.WriteFile("testdata/model.go", fmtSource, 0644); err != nil {
		log.Fatal(err)
	}
}
