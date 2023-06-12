package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches = []string{}
	wg      = sync.WaitGroup{}
	lock    = sync.Mutex{}
)

func Searching(root, filename string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(1)
	}
	for _, file := range files {
		fmt.Println(filepath.Join(root, file.Name()))
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go Searching(filepath.Join(root, file.Name()), filename)
		}
	}
	wg.Done()

}

func main() {

	wg.Add(1)
	go Searching("/home/kshitij/Desktop", "README.md")
	wg.Wait()

	fmt.Println(matches)
}
