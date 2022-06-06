package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

type UserRecord struct {
	UserId string
	Count  map[string]int
}

func main() {

	str := "ALTER TABLE order_info_%d DROP INDEX idx_username;\n"
	for i := 0; i < 128; i++ {
		fmt.Printf(str, i)
	}
	open, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewReader(open)
	regex := regexp.MustCompile(".+userId=(.+), countInfo=(.+)")

	users := []UserRecord{}

	for {
		line, err := buf.ReadString('\n')

		if err == io.EOF {
			break
		}

		group := regex.FindAllStringSubmatch(line, -1)
		userId := group[0][1:2]
		countInfo := group[0][2:3]

		m := make(map[string]int, 1)

		json.Unmarshal([]byte(countInfo[0]), &m)
		users = append(users, UserRecord{userId[0], m})
		//fmt.Println(userId[0], countInfo[0])
	}

	//url := "http://10.180.26.114:8090/backdoor/updateUserRecord?username=%s&productId=%s&amount=%d"
	for _, user := range users {
		fmt.Println(user.UserId)

		for k, v := range user.Count {
			if k != "" && v != 0 {
				//furl := fmt.Sprintf(url, user.UserId, k, v)
				//fmt.Println(furl)
				//http.Get(furl)
			}
		}
	}
}
