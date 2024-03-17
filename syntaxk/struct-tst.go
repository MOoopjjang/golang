package main

import (
	"fmt"
)

type Advertiser struct {
	Id   int64
	Name string
	Age  int32
	Addr string
}

func (advertiser *Advertiser) ToString() string {
	return fmt.Sprintf("id = %d , name = %s , age = %d , addr = %s", (*advertiser).Id, (*advertiser).Name, (*advertiser).Age, (*advertiser).Addr)
}

func main() {
	advertiser := Advertiser{1, "cwkim", 47, "incheon"}
	fmt.Println(advertiser.ToString())
}
