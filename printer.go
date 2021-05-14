package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd.exe", "/k", "explorer.exe 控制面板\\硬件和声音\\设备和打印机")

	output, err := cmd.Output()

	if err != nil {
		//panic(err.Error())
	}

	fmt.Println(string(output))

	open, err := os.Open("C:\\Users\\shidt\\Desktop\\硬件和声音.lnk")

	if err != nil {
		panic(err)
	}

	fmt.Println(ioutil.ReadAll(open))

	arr := [...]int{10, 12, 13}

	fmt.Println(arr)
}
