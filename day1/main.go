package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input *bufio.Scanner) int {
	calibrationValue := 0
	var err error
	for input.Scan() {
		// fmt.Println(s.Text())
		currLine := input.Text()
		l := 0
		r := len(currLine) - 1

		var foundLeft, foundRight bool
		var leftInt, rightInt int
		for l <= r {
			if !foundLeft {
				leftInt, err = strconv.Atoi(string(currLine[l]))
				if err == nil {
					foundLeft = true
				}
			}
			if !foundRight {
				rightInt, err = strconv.Atoi(string(currLine[r]))
				if err == nil {
					foundRight = true
				}

			}
			if !foundLeft {
				l++
			}
			if !foundRight {
				r--
			}
			if foundLeft && foundRight {
				break
			}
		}
		calibrationValue += leftInt*10 + rightInt
		// break
	}
	return calibrationValue
}

func rev(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

func part2(input *bufio.Scanner) int {
	calibrationValue := 0
	for input.Scan() {
		// fmt.Println(s.Text())
		digits := map[string]int{
			"zero":  0,
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}
		currLine := input.Text()

		var leftInt, rightInt int
		// find left
		foundLeft := false
		for i := 0; i < len(currLine); i++ {
			// find by string
			for k, _ := range digits {
				if strings.HasPrefix(string(currLine[i:]), k) {
					leftInt = digits[k]
					foundLeft = true
					break
				}
			}
			if foundLeft {
				break
			}
			// find by symbol
			p, err := strconv.Atoi(string(currLine[i]))
			if err == nil {
				leftInt = p
				break
			}
		}

		// find right
		foundRight := false
		for i := len(currLine) - 1; i >= 0; i-- {
			for k, _ := range digits {
				if strings.HasSuffix(string(currLine[:i+1]), k) {
					rightInt = digits[k]
					foundRight = true
					break
				}
			}
			if foundRight {
				break
			}
			// find by symbol
			p, err := strconv.Atoi(string(currLine[i]))
			if err == nil {
				rightInt = p
				break
			}
		}

		calibrationValue += leftInt*10 + rightInt
		// break
	}
	return calibrationValue
}

func main() {
	// read file
	fptr := flag.String("in", "in.txt", "./test.txt")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	calibrationValue := part2(s)

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer is ", calibrationValue)
}
