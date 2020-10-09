package main

import "fmt"

type request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *request, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		case req := <-service:
			go run(op, req)
		}
	}
}

func startServer(op binOp) (service chan *request, quit chan bool) {
	service = make(chan *request)
	quit = make(chan bool)
	go server(op, service, quit)

	return service, quit
}

func main() {
	service, quit := startServer(func(a, b int) int {
		return a + b
	})

	const N = 100

	var reqs [N]request

	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		service <- req
	}

	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("request ", i, " is ok")
		}
	}

	quit <- true
	fmt.Println("done")
}
