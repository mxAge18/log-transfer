package conf

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf `ini:"etcd"`
	ESConf `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
	ChanMaxSize int `ini:"chann_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	TimeOut time.Duration `ini:"timeout"`
	TaillogKey string `ini:"log_agent_key"`
}

type ESConf struct {
	Address string `ini:"address"`
	ChanSize int `ini:"chan_size"`
}

func InitConf(path string) (*AppConf) {
	var cfg AppConf
	err := ini.MapTo(&cfg, path)
	if err != nil {
		log.Fatalln("load config.ini fail, err=", err)
	}
	log.Println("config load success")
	return &cfg
}