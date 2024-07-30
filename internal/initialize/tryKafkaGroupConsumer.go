package initialize

import (
	"context"
	"gproject/internal/initialize/global"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var producer1 sarama.SyncProducer

func TryKafka1() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	producer1, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		global.Logger.Error("Failed to start Sarama producer")
	}
	defer func() {
		if err := producer1.Close(); err != nil {
			global.Logger.Error("Failed to close Sarama producer")
		}
	}()

	// Create 3 partitions for the "test-topic"
	admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, config)
	if err != nil {
		global.Logger.Error("Failed to create cluster admin")
	}
	defer func() {
		if err := admin.Close(); err != nil {
			global.Logger.Error("Failed to close cluster admin")
		}
	}()

	client, _ := sarama.NewClient([]string{"localhost:9092"}, config)
	partitions, _ := client.Partitions("test-topic")

	if len(partitions) < 3 {
		err = admin.CreatePartitions("test-topic", 3, nil, false)
		if err != nil {
			global.Logger.Error("Failed to create partitions", zap.Error(err))
		}
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/produce", produceMessage1)
	go func() {
		if err := r.Run(":8080"); err != nil {
			global.Logger.Error("Failed to start Gin server")
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 1; i <= 3; i++ {
		buildConsumer1(ctx, i)
	}
}

func buildConsumer1(ctx context.Context, consumerID int) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky

	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-consumer-group", config)
	if err != nil {
		global.Logger.Error("Failed to start Sarama consumer group", zap.Int("ConsumerID", consumerID))
		panic("")
	}
	defer consumer.Close()

	topics := []string{"test-topic"}
	handler := &ConsumerGroupHandler{consumerID: consumerID}

	for {
		err := consumer.Consume(ctx, topics, handler)
		if err != nil {
			global.Logger.Error("Error from consumer", zap.Int("ConsumerID", consumerID), zap.Error(err))
		}
	}
}

type ConsumerGroupHandler struct {
	consumerID int
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	global.Logger.Info("Consumer ready", zap.Int("ConsumerID", h.consumerID))
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	global.Logger.Info("Consumer closing", zap.Int("ConsumerID", h.consumerID))
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		global.Logger.Info("Message claimed",
		 zap.Int("ConsumerID", h.consumerID), 
		 zap.String("Value", string(msg.Value)),
		)
		session.MarkMessage(msg, "")
	}
	return nil
}

func produceMessage1(c *gin.Context) {
	message := c.Query("message")
	topic := "test-topic"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := producer1.SendMessage(msg)
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
		"number_partitions": len(partitions),
	})
}
