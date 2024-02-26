package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	open, err := os.Open("/Users/star/go_ws/LearnGo/regex/data.txt")

	fmt.Print(err)
	reader := bufio.NewReader(open)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		cmd := string(line)
		fmt.Println(cmd)
		// #XMDT#{__traceId__=8925812202412258150}#XMDT#
		regex := regexp.MustCompile(".+#ppp#\\{__traceId__=(.+)\\}.+")
		match := regex.MatchString(cmd)

		if match {
			submatch := regex.FindAllStringSubmatch(cmd, -1)
			fmt.Println(submatch[0][0])
		}
	}
}
