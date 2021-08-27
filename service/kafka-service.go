package service

import (
	"fmt"
	"log"
	"log-transfer/entity"

	"gopkg.in/Shopify/sarama.v1"
)
type KafkaService interface {
	ConsumeData()
	initService()
}

func NewKafka(address []string, topic string, es ElasticService) KafkaService {
	kafka := kafkaService{
		address: address,
		topic: topic,
		config: sarama.NewConfig(),
		es: es,
	}
	kafka.initService()
	return &kafka
}

type kafkaService struct {
	Consumer sarama.Consumer
	partitionList []int32
	address []string
	topic string
	config *sarama.Config
	err error
	es ElasticService
}

func (this *kafkaService) initService() {
	this.Consumer, this.err = sarama.NewConsumer(this.address, this.config)
	if this.err != nil {
		log.Fatalln("new consumer fail, err", this.err)
	}
	this.partitionList, this.err = this.Consumer.Partitions(this.topic)
	if this.err != nil {
		log.Fatalln("consumer get partitions fail, err", this.err)
	}
	log.Println("分区：", this.partitionList)
	log.Println("Topic", this.topic)
}
func (this *kafkaService) ConsumeData() {
	for partition := range this.partitionList {
		// 针对每个分区创建分区消费者
		pc, err := this.Consumer.ConsumePartition(this.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Println("consumer get ConsumePartition fail, err", err)
			return
		}
		// defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d, offset %d, key %v, value %v\n", msg.Partition,
			msg.Offset,msg.Key, string(msg.Value))
				var data entity.ESData
				data.LogDetail = string(msg.Value)
				data.Topic = msg.Topic
				this.es.SendDataToESChan(&data)
			}
		}(pc)
	}
}


