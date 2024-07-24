package initialize

import (
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

	var err error
	producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama producer:", err)
		}
	}()

	// Initialize Gin router
	r := gin.Default()

	r.POST("/send-message", func(c *gin.Context) {
		message := c.Query("message")

		msg := &sarama.ProducerMessage{
			Topic: "test-topic",
			Value: sarama.StringEncoder(message),
		}

		_, _, err := producer.SendMessage(msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	})

	// Start Gin server in a Goroutine
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to start Gin server: %v", err)
		}
	}()

	// Initialize Kafka consumer
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Failed to start Sarama partition consumer:", err)
	}

	// Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d: %s\n", msg.Offset, string(msg.Value))
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Println("Consumer shutting down...")

}
