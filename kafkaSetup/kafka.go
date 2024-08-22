package repository

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

var (
	defaultKafkaProducer sarama.SyncProducer
	defaultKafkaConsumer sarama.Consumer
)

type Kafka struct {
	Host     string `mapstructure:"host"`
	Port     int  `mapstructure:"port"`
}

func InitKafka() {
	kafkaCfg := Kafka{
		Host: viper.GetString("kafka.host"),
		Port: viper.GetInt("kafka.port"),
	}
	var err error
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	saramaConfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	//saramaConfig.Consumer.Offsets.AutoCommit.Enable = true
	addr := fmt.Sprintf("%s:%d", kafkaCfg.Host, kafkaCfg.Port)

	// 连接kafka, 初始化生产者
	producer, err := sarama.NewSyncProducer([]string{addr}, saramaConfig)
	if err != nil {
		log.Fatalf("producer closed, err:%v", err)
		return
	}
	defaultKafkaProducer = producer

	// 连接kafka, 初始化消费者
	consumer, err := sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		log.Fatalf("fail to start consumer, err:%v\n", err)
	}
	defaultKafkaConsumer = consumer
}

func GetKafkaProducer() sarama.SyncProducer {
	return defaultKafkaProducer
}

func GetKafkaConsumer() sarama.Consumer {
	return defaultKafkaConsumer
}
