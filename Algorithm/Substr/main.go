package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	bTarget, _, err := scanner.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	target := string(bTarget)

	bSep, _, err := scanner.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	sep := string(bSep)

	var result string

	for {
		if !strings.Contains(target, sep) {
			result += target
			break
		}

		before, after, _ := strings.Cut(target, sep)
		result = result + before
		target = after
	}

	log.Println(result)
}