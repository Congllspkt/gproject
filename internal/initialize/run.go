package initialize

import (
	"gproject/internal/initialize/global"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

var producer sarama.SyncProducer

func Run() {

	InitConFig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	TryDataSample()

	// Initialize Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	producer, _ = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	defer func() {
		producer.Close()
	}()

	// Initialize Gin router
	r := gin.Default()
	r.POST("/send-message", func(c *gin.Context) {
		message := c.Query("message")
		msg := &sarama.ProducerMessage{
			Topic: "test-topic",
			Value: sarama.StringEncoder(message),
		}
		producer.SendMessage(msg)
		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	})

	// Start Gin server in a Goroutine
	go func() {
		r.Run(":8080")
	}()
	// Initialize Kafka consumer
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, config)
	partitionConsumer, _ := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			global.Logger.Info("Consumed message offset " + string(msg.Offset) + ": " + string(msg.Value) + "\n")
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Println("Consumer shutting down...")

}
