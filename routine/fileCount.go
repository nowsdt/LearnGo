package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"sync"
)

var sema sync.WaitGroup
var ch = make(chan CountInfo, 100)

type CountInfo struct {
	name string
	line int
}

type CountInfos []CountInfo

func (x CountInfos) Len() int {
	return len(x)
}

func (x CountInfos) Less(i, j int) bool {
	return x[i].line > x[j].line
}
func (x CountInfos) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	log.Println("start")
	pathDir := "/Users/star/"
	sema.Add(1)
	go walkDir(pathDir)

	counts := make([]CountInfo, 100)

	go func() {
		for data := range ch {
			//log.Println(data)
			counts = append(counts, data)
		}
	}()

	sema.Wait()
	//time.Sleep(10 * time.Second)
	log.Println("end")

	sort.Sort(CountInfos(counts))
	for i, count := range counts {
		if count.name == "" {
			continue
		}
		if i >= 50 {
			break
		}
		log.Println(i, count.name, count.line)
	}

}

func walkDir(pathDir string) {
	defer sema.Done()
	//log.Println("pathDir:", pathDir)

	dir, err := ioutil.ReadDir(pathDir)

	if err != nil {
		log.Printf("pathDir:%s read error %s\n", pathDir, err)
		return
	}

	for _, info := range dir {
		if info.IsDir() {
			sema.Add(1)

			//walkDir(path.Join(pathDir, info.Name()))
			go walkDir(path.Join(pathDir, info.Name()))
		} else {
			if len(info.Name()) == 0 || path.Ext(info.Name()) != ".java" {
				continue
			}
			count, _ := lineCount(path.Join(pathDir, info.Name()))
			ch <- CountInfo{info.Name(), count}
		}
	}
}

func lineCount(filePath string) (line int, err error) {
	//log.Println("filePath:", filePath)

	var readSize int

	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("%s open err %s \n", filePath, err)
		return 0, err
	}

	defer f.Close()
	buf := make([]byte, 1024)
	for {
		readSize, err = f.Read(buf)
		if err != nil {
			//log.Println(err)
			break
		}
		var bufPos int
		for {
			i := bytes.IndexByte(buf[bufPos:], '\n')
			if i == -1 || readSize == bufPos {
				break
			}
			bufPos = bufPos + i + 1
			line++
		}

	}

	if readSize > 0 && line == 0 || line > 0 {
		line++
	}
	if err == io.EOF {
		return line, nil
	}
	return line, err

}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
