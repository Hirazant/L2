package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func cd(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		return
	}
}

func kill(id int) {
	pr, err := os.FindProcess(id)
	if err != nil {
		log.Fatal(err)
	}
	err = pr.Kill()
	if err != nil {
		log.Fatal(err)
	}
}

func ps() {
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(os.Stdout, "%s$: ", path)

		input, err := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		command := strings.Split(input, " ")

		switch command[0] {
		case "\\quit":
			break
		case "cd":
			if len(command) == 2 {
				cd(command[1])
			} else {
				fmt.Fprintf(os.Stdout, "incorrect input\n")
			}
		case "pwd":
			fmt.Fprintf(os.Stdout, "%s\n", path)
		case "echo":
			if len(command) == 2 {
				fmt.Fprintf(os.Stdout, command[1])
			} else {
				fmt.Fprintf(os.Stdout, "incorrect input\n")
			}
		case "kill":
			id, err := strconv.Atoi(command[1])
			if err != nil {
				log.Fatal(err)
			}
			kill(id)
		case "ps":
			ps()
		}

	}
}
