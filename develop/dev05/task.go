package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res, nil
}

func readFromConsole() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			// exit if user entered an empty string
			break
		}
	}
	return input
}

func match(str string, pat string) (bool, error) {
	return regexp.MatchString(pat, str)
}

func afterPrint(lines []string, index int, after int) {
	if index+after+1 < len(lines) {
		for i := index + 1; i <= index+after; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := index + 1; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
}

func beforePrint(lines []string, index int, before int) {
	if index-before >= 0 {
		for i := index - before; i < index; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := 0; i < index; i++ {
			fmt.Println(lines[i])
		}
	}
}

func grep(after int, before int, count bool, invert bool, lineNum bool, pattern string, fileName string) {
	var lines []string
	var err error

	if fileName == "-" {
		lines = readFromConsole()
	} else {
		lines, err = readFromFile(fileName)
	}

	if err != nil {
		log.Fatalf("cannot open file")
	}
	var cou int

	for index, val := range lines {

		ok, err := match(val, pattern)
		if err != nil {
			log.Fatalf("error with matchking")
		}

		// Инвертируем, если передан соотв. флаг
		ok = ok != invert

		if ok == true {

			cou++

			if lineNum {
				fmt.Println(index)
			} else {
				if before > 0 {
					beforePrint(lines, index, before)
				}

				fmt.Println(val)

				if after > 0 {
					afterPrint(lines, index, after)
				}
			}
		}
	}

	if count {
		fmt.Printf("found in %d lines\n", cou)
	}

}

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatalf("File name or - for console input and pattern required")
	}

	args := flag.Args()
	pattern := args[0]
	fileName := args[1]

	// Если передан контекст, то переопределяем в афтер и бефор
	if *context != 0 {
		after = context
		before = context
	}

	// Изменяем паттерн, если передан соотв. флаг
	if *ignoreCase {
		pattern = "(?i)" + pattern
	}

	// Редактируем паттерн на поиск целой строки
	if *fixed {
		pattern = "^" + pattern + "$"
	}

	grep(*after, *before, *count, *invert, *lineNum, pattern, fileName)
}
