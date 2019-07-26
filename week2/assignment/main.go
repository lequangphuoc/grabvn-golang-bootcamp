package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func readFileByLine(filePath string, lines chan<-string, processingFile <-chan string, wg *sync.WaitGroup) {	
	file, err := os.Open(filePath)
	
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		file.Close();
		
		<- processingFile
		if len(processingFile) == 0 {
			close(lines)
		}

		wg.Done()
    }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()		
	}
}

func countWords(lines chan string, wordCount map[string]int, wg *sync.WaitGroup)  {
	for line := range lines {
		words := strings.Fields(line)

		for _, word := range words { 
			wordCount[word]++
		}
	}

	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	path := "files"
	
	files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
	}
	
	lines := make(chan string, len(files))
	processingFile := make(chan string, len(files))
	wordCount := make(map[string]int)

	for _, file := range files {
		processingFile <- file.Name()
	}

    for _, file := range files {
		filePath := path + "/" + file.Name()
		wg.Add(1)
		go readFileByLine(filePath, lines, processingFile, &wg)
	}

	wg.Add(1)
	go countWords(lines, wordCount, &wg)
	
	wg.Wait()

	for key, val := range wordCount {
		fmt.Println(key, ": ", val)
	}
}

