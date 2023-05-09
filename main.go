package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/raphael-foliveira/goYtDownloader/downloader"
)

var waitGroup sync.WaitGroup

func main() {
	err := os.RemoveAll("./downloads")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("./downloads", 0755)
	if err != nil {
		fmt.Println("Directory already exists")
	}
	args := os.Args[1:]
	downloaders := []*downloader.Downloader{}
	for _, arg := range args {
		downloaders = append(downloaders, &downloader.Downloader{VideoId: arg})
	}
	for _, d := range downloaders {
		waitGroup.Add(1)
		go d.Download(&waitGroup)
	}
	waitGroup.Wait()
}
