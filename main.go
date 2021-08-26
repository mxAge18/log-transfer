package main

import (
	"fmt"
	"log-transfer/conf"
)



func main() {
	cfg := conf.InitConf("./conf/config.ini")
	fmt.Println(cfg.EtcdConf.Address)
}