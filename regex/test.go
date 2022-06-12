package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	s := "133alhgalghg userId=155 json={JN}"
	regex := regexp.MustCompile(".{1,12} userId=(.+) json=(.+)")
	match := regex.MatchString(s)

	if match {
		submatch := regex.FindAllStringSubmatch(s, -1)

		//fmt.Println(submatch)
		fmt.Println(submatch[0][0:1])
		fmt.Println(submatch[0][1:2])
		fmt.Println(submatch[0][2:3])
	}
	open, err := os.Open("P:\\多媒体整理\\20220105\\remove\\新建文件夹\\20220105.txt")

	fmt.Print(err)
	reader := bufio.NewReader(open)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		cmd := string(line)
		fmt.Println(cmd)

		command := exec.Command(cmd)
		err = command.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}
