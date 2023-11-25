package queue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

// ProducerInterface defines the interface for a Kafka producer.
type ProducerInterface interface {
	// SendMsg sends a message to the specified Kafka topic using the provided input.
	SendMsg(QProducerInput) *wraperror.WrappedError
}

// QProducerInput represents the input structure for sending messages to Kafka.
type QProducerInput struct {
	Service string      // Kafka topic to send the message to
	Message interface{} // Message content
}

// Producer is a struct implementing the ProducerInterface.
type Producer struct {
	QProducer *kafka.Producer
}

// NewProducer creates and returns a new Kafka producer based on the provided configuration.
func NewProducer(conf config.InterfaceConfig) (ProducerInterface, *wraperror.WrappedError) {
	kafkaconfig := kafka.ConfigMap{
		"bootstrap.servers": conf.GetString("queue.host"),
		"group.id":          conf.GetString("queue.group_id"),
		"auto.offset.reset": "earliest",
	}
	producer, err := kafka.NewProducer(&kafkaconfig)
	if err != nil {
		werr := wraperror.Wrap(err, "error encountered during new producer creation", "error", http.StatusInternalServerError)
		return nil, werr
	}

	return &Producer{QProducer: producer}, nil
}

// SendMsg sends a message to the specified Kafka topic using the provided input.
func (producer *Producer) SendMsg(input QProducerInput) *wraperror.WrappedError {
	// Produce messages to the topic (asynchronously)

	// Convert the message to JSON
	inputMsg, _ := json.Marshal(input.Message)

	// Create a delivery channel for Kafka events
	deliveryChan := make(chan kafka.Event)

	// Produce the message to the Kafka topic
	err := producer.QProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &input.Service, Partition: kafka.PartitionAny},
		Value:          inputMsg,
	}, deliveryChan)

	if err != nil {
		return wraperror.Wrap(err, "error producing message", "error", http.StatusInternalServerError)
	}

	// Wait for the delivery report
	e := <-deliveryChan
	m := e.(*kafka.Message)

	// Check if the delivery was successful
	if m.TopicPartition.Error != nil {
		return wraperror.Wrap(m.TopicPartition.Error, "delivery failed", "error", http.StatusInternalServerError)
	}

	// Print information about the delivered message
	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	// Close the delivery channel
	close(deliveryChan)

	return nil
}
