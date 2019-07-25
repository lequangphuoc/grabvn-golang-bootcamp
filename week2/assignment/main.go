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

func count(filePath string, wordCount map[string]int, waitgroup *sync.WaitGroup) {
	
	f, err := os.Open(filePath)
	
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close();
		waitgroup.Done()
    }()

	s := bufio.NewScanner(f)
	for s.Scan() {
		words := strings.Fields(s.Text())

		for _, word := range words {			
			wordCount[word]++
		}			
	}
}

func main() {
	var waitgroup sync.WaitGroup
	path := "files"
	files, err := ioutil.ReadDir(path)
	//done := make(chan string)
    if err != nil {
        log.Fatal(err)
	}
	
	wordCount := make(map[string]int)

    for _, file := range files {
		filePath := path + "/" + file.Name()
		waitgroup.Add(1)
		go count(filePath, wordCount, &waitgroup)
	}
	
	waitgroup.Wait()
	for key, val := range wordCount {
		fmt.Println(key, ": ", val)
	}
}

