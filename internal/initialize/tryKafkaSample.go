package initialize

import (
	"gproject/internal/initialize/global"
	"net/http"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var producer sarama.SyncProducer

func TryKafka() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		global.Logger.Error("Failed to start Sarama producer")
	}
	defer func() {
		if err := producer.Close(); err != nil {
			global.Logger.Error("Failed to close Sarama producer")
		}
	}()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/produce", produceMessage)
	go func() {
		if err := r.Run(":8080"); err != nil {
			global.Logger.Error("Failed to start Gin server")
		}
	}()

	buildConsumer()
}

func buildConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		global.Logger.Error("Failed to start Sarama consumer")
		panic("")
	}

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		global.Logger.Error("Failed to start Sarama partition consumer")
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			global.Logger.Info("Consumed message offset:",
				zap.Int64("Offset", msg.Offset),
				zap.String("Value", string(msg.Value)),
				zap.String("Partition", string(msg.Partition)),
			)
		case <-signals:
			break ConsumerLoop
		}
	}
	global.Logger.Info("Consumer shutting down...")
	panic("done")
}

func produceMessage(c *gin.Context) {
	message := c.Query("message")
	topic := "test-topic"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	broker := "localhost:9092"
	config := sarama.NewConfig()
	client, _ := sarama.NewClient([]string{broker}, config)
	partitions, _ := client.Partitions("test-topic")
	c.JSON(http.StatusOK, gin.H{
		"message":           "Message sent successfully",
		"topic":             topic,
		"number partitions": len(partitions),
	})
}
