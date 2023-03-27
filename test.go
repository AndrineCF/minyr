
package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	//Open fil
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")

	//check if there was error when open the file
	if err != nil {
		log.Fatal(err)
	}
	
	//closing the fill when finished
	defer src.Close()

	log.Println(src)

	//Scaning the file
	scanner := bufio.NewScanner(src)

	// variable to stored the scanner text
	var lines []string
	
	// go through the file
	for scanner.Scan() {
		lineSpilt := strings.Split(scanner.Text(), ";")
		if len(lineSpilt) > 3 {
			log.Println(lineSpilt[3])
		}
		lines = append(lines, scanner.Text())
	}

	fmt.Println(lines[0])
	//checking scanner
	errScanner := scanner.Err()
	if errScanner != nil {
		log.Fatal(errScanner)
	}

}
