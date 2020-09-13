package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%0.2f", d)
}

func (db database) list(res http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(res, "%s: %s\n", item, price)

	}
}
func (db database) price(res http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if ok {
		fmt.Fprintf(res, "%s: %s\n", item, price)
	} else {
		fmt.Fprintf(res, "%s: not found\n", item)
	}
}

func (db database) add(res http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	fPrice, err := strconv.ParseFloat(price, 32)
	if err == nil {
		db[item] = dollars(fPrice)
		fmt.Fprint(res, "OK")
	}
}

func (db database) update(res http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	fPrice, err := strconv.ParseFloat(price, 32)
	if err == nil {
		_, ok := db[item]
		if ok {
			db[item] = dollars(fPrice)
			fmt.Fprint(res, "OK")
		} else {
			fmt.Fprint(res, "not exists")
		}
	}
}

func (db database) del(res http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if ok {
		delete(db, item)
		fmt.Fprint(res, "OK\n")
	} else {
		fmt.Fprint(res, "%s: not found\n", item)
	}
}

func server1() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/add", db.add)
	mux.HandleFunc("/del", db.del)
	log.Fatal(http.ListenAndServe("localhost:8001", mux))
	// 使用默认的 DefaultServeMux
	//log.Fatal(http.ListenAndServe("localhost:8001", nil));
}

func (db database) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(res, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if ok {
			fmt.Fprintf(res, "%s: %s\n", item, price)
		} else {
			fmt.Fprintf(res, "%s: not found\n", item)
		}
	default:
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "no such page,%s\n", req.URL)
	}
}

func server0() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func main() {
	//server0()
	server1()
}
