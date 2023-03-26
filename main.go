package main

import (
	"fmt"
	"os"
	"log"
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

	fmt.Println("test")
}
