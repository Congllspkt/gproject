package initialize

import (
	"context"
	"gproject/internal/initialize/global"
	"math/rand/v2"
	"net/http"

	// "time"

	"github.com/IBM/sarama"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var producer1 sarama.SyncProducer

var visited = make(map[string]int)

func TryKafka1() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewManualPartitioner

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

	go buildConsumer1(ctx, 0)
	go buildConsumer1(ctx, 1)
	buildConsumer1(ctx, 2)
}

func buildConsumer1(ctx context.Context, consumerID int) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, _ := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-consumer-group", config)
	defer consumer.Close()

	topics := []string{"test-topic"}
	handler := &ConsumerGroupHandler{consumerID: consumerID}
	for {
		consumer.Consume(ctx, topics, handler)
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
	// time.Sleep(5 * time.Second)
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if msg.Value[0] == 'A' {
			global.LoggerConsumer.Info("receive - fail",
				zap.Int("ConsumerID", h.consumerID),
				zap.String("Value", string(msg.Value)),
			)
			key := string(msg.Value)
			if val, ok := visited[key]; ok {
				visited[key] = val + 1
			} else {
				visited[key] = 1
			}

			if visited[key] >= 3 {
				global.LoggerConsumer.Info("Consumer read 3 times, get Error",
					zap.Int("ConsumerID", h.consumerID),
					zap.String("Value", string(msg.Value)),
				)
				delete(visited, key)
				session.MarkMessage(msg, "")
			} 
			return nil
		}
			global.LoggerConsumer.Info("received",
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

	num := rand.IntN(3)
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(num),
		Value:     sarama.StringEncoder(message),
	}
	_, _, err := producer1.SendMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	global.LoggerProducer.Info("Sent",
		zap.String("message", message),
		zap.String("topic", topic),
		zap.Int("partition", num),
	)

	c.JSON(http.StatusOK, gin.H{
		"message":           message,
		"topic":             topic,
		"partition":         num,
	})
}
