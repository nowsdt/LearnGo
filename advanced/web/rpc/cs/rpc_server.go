// rpc_server.go
// after client-exits the server shows the message:
//       1:1234: The specified network name is no longer available.
//       2011/08/01 16:19:04 rpc: rpc: server cannot decode request: WSARecv tcp 127.0.0.
package main

//  shidt 2020.10.12 10:50
//go一个目录下面不能有多个包packgae，因此把rpc_server和rpc_client单独拿出来
import (
	"code.com/advanced/web/rpc"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	log.Println("rpc server started")
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}

/* Output:
Starting Process E:/Go/GoBoek/code_examples/chapter_14/rpc_server.exe ...

** after 5 s: **
End Process exit status 0
*/
