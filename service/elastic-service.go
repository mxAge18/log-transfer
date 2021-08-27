package service

import (
	"context"
	"encoding/json"
	"log"
	"log-transfer/entity"
	"time"

	"github.com/olivere/elastic/v7"
)
type ElasticService interface {
	initService()
	SendDataToES()
	SendDataToESChan(data *entity.ESData)
}

func NewElastic(addr string, chanSize int) ElasticService {
	elastic := elasticService{
		address: addr,
		chanSize: chanSize,
	}
	elastic.initService()
	go elastic.SendDataToES()
	return &elastic
}

type elasticService struct{
	ESClient *elastic.Client
	err error
	address string
	chanSize int
	ESDataChan chan *entity.ESData
}

func (this *elasticService) initService() {
	log.Println(this.address)
	sniffOption := elastic.SetSniff(false)
	this.ESClient, this.err = elastic.NewClient(elastic.SetURL(this.address), sniffOption)
	this.ESDataChan = make(chan *entity.ESData, this.chanSize)
	if this.err != nil {
		log.Fatalln("es conncetion fail ,err=", this.err)
	}
	log.Println("es connection success")
}
func (this *elasticService) SendDataToESChan(data *entity.ESData) {
	this.ESDataChan <- data
}

func (this *elasticService) SendDataToES() {
	for {
		select{
			case msg := <- this.ESDataChan:
				index := msg.Topic
				data, err := json.Marshal(msg)
				if err != nil {
					log.Fatalln("data marshal fail, err=", err)
				}
				put, err := this.ESClient.Index().Index(index).Type("_doc").BodyJson(string(data)).Do(context.Background())
		
				if err != nil {
					log.Fatalln("put data to es fail, err=", err)
				}
				log.Println("put data to es succss, index=", put.Index)
			default:
				time.Sleep(time.Millisecond * 50)
		}
	}


}

