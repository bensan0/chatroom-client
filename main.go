package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn,err:=net.Dial("tcp", "127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	//獲取伺服端訊息
	go sendMsg(os.Stdout, conn)
	
	//將用戶輸入的訊息傳送到伺服端
	sendMsg(conn, os.Stdin)
}

func sendMsg(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
