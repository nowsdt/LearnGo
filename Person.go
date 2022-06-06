package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func main() {

	//DROP INDEX idx_createtime_producttype,
	//DROP INDEX idx_origin_orderid,
	//DROP INDEX idx_username

	str := "ALTER TABLE order_info_%d DROP INDEX idx_username,DROP INDEX idx_origin_orderid, DROP INDEX idx_createtime_producttype;\n"
	//str2 := "ALTER TABLE xproduct_ext_info_%d DROP INDEX idx_productcode,DROP INDEX idx_abnormal_reason;\n"
	for i := 0; i < 128; i++ {
		//fmt.Printf(str, i)
		fmt.Printf(str, i)
	}
}
