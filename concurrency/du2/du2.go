package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

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
	// FIXME
	//fmt.Println("try to get")
	//sema <- struct{}{}  //多这一行会导致程序进行不下去
	//fmt.Println("success got")
	//fmt.Println("got")
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

func main() {
	/*	flag.Parse()
		roots := flag.Args()*/
	roots := os.Args[1:]

	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		read, _ := os.Stdin.Read(make([]byte, 1))
		fmt.Println("read:", read)
		close(done)
	}()

	/*	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("sem.len:", len(sema))
		}
	}()*/

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) // 一个walkDir 遍历dD: 39s
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)

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
