package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Know-more Daemon, version - v1.0.0")
	fmt.Println("Know-more is a knowledge base that stores knowledge and intercepts it live from bash shells ")

	file, err := os.Open("/tmp/trace")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
