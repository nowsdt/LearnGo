package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	if canceled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v \n", err)
		return nil
	}

	return entries
}

var verbose = flag.Bool("v", false, "show berbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	startTime := time.Now().UnixNano()
	var n sync.WaitGroup
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDir(root, &n, fileSizes) // 一个walkDir 遍历dD: 39s
		}
	}()

	go func() {
		n.Wait()
		endTime := time.Now().UnixNano()
		fmt.Printf("time escape:%0.1fs\n", float64((endTime-startTime)/1e9))
		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nFiles, nBytes int64

loop:
	for {
		select {
		case <-done:
			for range fileSizes {

			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSize 关闭
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}

	/*	for i := range fileSizes {
		nFiles = nFiles + i
		nBytes++
	}*/

	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nFiles int64, total int64) (int, error) {
	return fmt.Printf("%d files, size: %0.1f GB\n", nFiles, float64(total)/1e9)
}

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
