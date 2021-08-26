package service

type KafkaService interface {
	GetDataFromTopic()
	SendDataToESChan()
}

type ElasticService interface {
	SendDataToES()
	CreateIndex()
}

type LogTransService interface {

}