package main

import (
	"fmt"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	dbPath string = "./ip2region.xdb"
	cBuff  []byte
)

func init() {
	var err error
	cBuff, err = xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}
}

func main() {
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	var ip = "120.229.81.7"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Printf("region: %s, took: %s\n", region, time.Since(tStart))
}
