package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

func InitKafka() {

}

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

// producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s
	jsonBody, _ := json.Marshal(body)

	// create msg
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)

	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action success",
	})
}

func RegisterConsumeATC (id int) {
	kafkaGroupId := "consumer_group-"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	fmt.Printf("Brokers (%v) \n", reader.Config().Brokers)

	

	defer reader.Close()

	fmt.Printf("Consumer (%d) Hong Phien ATC-----\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v  \n", id, err)
			return
		}

		fmt.Printf("Consumer (%d)\n", id)
		fmt.Printf("Topic (%v)\n", m.Topic)
		fmt.Printf("Partition (%v)\n", m.Partition)
		fmt.Printf("Offset (%v)\n", m.Offset)
		fmt.Printf("time (%d)\n", m.Time.Unix())
		fmt.Printf("time (%s)\n", string(m.Key))
		fmt.Printf("time (%s)\n", string(m.Value))
		

	}
}
