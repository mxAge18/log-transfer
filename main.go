package main

import (
	"log-transfer/conf"
	"log-transfer/service"
)

func main() {
	// 配置加载
	cfg := conf.InitConf("./conf/config.ini")
	// kafka 初始化
	// es 初始化
	var (
		es service.ElasticService = service.NewElastic(cfg.ESConf.Address, cfg.ESConf.ChanSize)
		k service.KafkaService = service.NewKafka([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic, es)
	)
	// 运行服务
	k.ConsumeData()
	select{}
}