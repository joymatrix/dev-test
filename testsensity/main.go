package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

func main() {

	fmt.Println("hello world")

	file, err := os.Open("./sensitive_words_lines.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dataSource := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataSource = append(dataSource, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: dataSource,
	})
	if err != nil {
		panic(err)
	}

	filterText := "你好真美丽"
	filterManage := filter.NewDirtyManager(memStore)
	var delim rune = '*'
	result, err := filterManage.Filter().Filter(filterText, delim)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
