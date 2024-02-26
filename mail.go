package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"time"
)

func main() {

	// https://help.sina.com.cn/comquestiondetail/view/798/
	address := "smtp.sina.com:465"
	from := "*@sina.com"

	host, _, _ := net.SplitHostPort(address)

	auth := smtp.PlainAuth("", from, "*", host)
	to := "*@qq.com"

	header := make(map[string]string)

	from_ := mail.Address{"robot", from}

	header["From"] = from_.String()
	header["To"] = to
	header["Subject"] = "开机了"
	header["Content-Type"] = "text/html;charset=UTF-8"

	body := "开机了 " + time.Now().Format("2006-01-02 15:04:05")

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + body

	var smtpClient *smtp.Client

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	//创建一个tls链接
	if conn, err := tls.Dial("tcp", address, tlsConfig); err != nil {
		log.Println("tls.Dial: ", err)
		return
	} else {
		smtpClient, err = smtp.NewClient(conn, host)
	}
	var err error

	err = smtpClient.Auth(auth)
	errBreak(err)

	err = smtpClient.Mail(from)
	errBreak(err)

	err = smtpClient.Rcpt(to)
	errBreak(err)

	data, err := smtpClient.Data()
	errBreak(err)

	_, err = data.Write([]byte(message))

	errBreak(err)

	err = data.Close()

	errBreak(err)

	//err := smtp.SendMail(host+":"+port, auth, from, to, message)
	log.Println("mail send ok")

	time.Sleep(time.Second * 3)
}

func errBreak(err error) {
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 3)
		panic(err)
		//log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags)
}
