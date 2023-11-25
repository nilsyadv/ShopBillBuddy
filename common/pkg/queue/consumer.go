package queue

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

// ConsumerInterface defines the interface for a Kafka consumer.
type ConsumerInterface interface {
	// ReadMsg reads messages from the Kafka topic and sends them to the provided output channels.
	ReadMsg(*QConsumerOutput)
	// SubscribeTopics subscribes the consumer to the specified Kafka topics.
	SubscribeTopics(...string)
}

// QConsumerOutput represents the output channels for Kafka messages and errors.
type QConsumerOutput struct {
	Msg   chan string                  // Channel to send Kafka messages
	Error chan *wraperror.WrappedError // Channel to send errors
}

// QConsumer is a struct implementing the ConsumerInterface.
type QConsumer struct {
	QConsumer *kafka.Consumer
}

// NewConsumer creates and returns a new Kafka consumer based on the provided configuration.
func NewConsumer(conf config.InterfaceConfig) (ConsumerInterface, *wraperror.WrappedError) {
	kafkaconfig := kafka.ConfigMap{
		"bootstrap.servers": conf.GetString("queue.host"),
		"group.id":          conf.GetString("queue.group_id"),
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(&kafkaconfig)
	if err != nil {
		werr := wraperror.Wrap(err, "error encountered during new consumer creation", "error", http.StatusInternalServerError)
		return nil, werr
	}

	return &QConsumer{QConsumer: consumer}, nil
}

// SubscribeTopics subscribes the consumer to the specified Kafka topics.
func (qcons *QConsumer) SubscribeTopics(topics ...string) {
	qcons.QConsumer.SubscribeTopics(topics, nil)
}

// ReadMsg reads messages from the Kafka topic and sends them to the provided output channels.
func (qcons *QConsumer) ReadMsg(output *QConsumerOutput) {
	defer close(output.Msg)

	for {
		msg, err := qcons.QConsumer.ReadMessage(-1)
		if err == nil {
			output.Msg <- string(msg.Value)
		} else {
			if err.(kafka.Error).Code() == kafka.ErrAllBrokersDown {
				fmt.Println("All brokers are down.")
				output.Error <- wraperror.Wrap(err, "all brokers are down", "warning", http.StatusInternalServerError)
				continue
			}
			output.Error <- wraperror.Wrap(err, "unknown error encountered", "error", http.StatusInternalServerError)
			continue
		}
	}
}
