package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type html2 struct {
	Title, Body string
}

var ex = []string{"content.html", ".id", ".DS_Store"}
var set map[string]bool
var t *template.Template

const baseDir = "/Users/star/dev_env/spkg/设计模式之美"
const dir = "/Users/star/dev_env/spkg/设计模式之美/origin"

func main() {
	filepath.Walk(dir, process)
}

func process(paths string, info fs.FileInfo, err error) error {
	if info == nil {
		log.Fatal(paths)
	}
	ext := path.Ext(info.Name())
	if info.IsDir() || ext == ".id" || !strings.HasSuffix(filepath.Dir(paths), "origin") {
		return err
	}
	if _, ok := set[info.Name()]; ok {
		return err
	}
	if ext == ".html" {
		out, err := os.OpenFile(path.Join(baseDir, "/html", info.Name()), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			log.Printf("create file %s fail", info.Name(), err)
			return err
		}

		in, err := os.Open(paths)
		if err != nil {
			log.Printf("read file %s fail", info.Name(), err)
			return err
		}
		readBytes, err := ioutil.ReadAll(in)
		err = t.Execute(out, html2{info.Name(), string(readBytes)})

	}

	if err != nil {
		log.Println(err)
	}
	//log.Println(paths, "------",info.Name())
	return err
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	set = make(map[string]bool)
	for _, s := range ex {
		set[s] = true
	}

	t, _ = template.ParseFiles("/Users/star/go_ws/LearnGo/advanced/rw/temp.html")
}
