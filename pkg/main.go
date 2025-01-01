package main

import (
	"bufio"
	"fmt"
	"know-more/pkg/core"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Know-more Daemon, version - v1.0.0")
	fmt.Println("Know-more is a knowledge base that stores knowledge and intercepts it live from bash shells ")
	config := core.InitializeConfig()
	file, err := os.Open(config.CommandsSourceFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println("Going to parse text: " + scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
