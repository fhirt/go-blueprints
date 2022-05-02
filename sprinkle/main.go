package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + " app",
	otherWord + " site",
	otherWord + " time",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
	otherWord + " hq",
}

func main() {
	path := flag.String("path", "", "the path to the file where additional transforms are defined")
	flag.Parse()
	customTransforms, err := readLines(*path)
	if err == nil {
		transforms = append(transforms, customTransforms...)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner((os.Stdin))
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
