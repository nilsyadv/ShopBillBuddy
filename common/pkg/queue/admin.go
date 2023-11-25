package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

type Topic struct {
	TopicName         string
	NumberOfPartition int
	ReplicationFactor int
}

func QueueSetup(conf config.InterfaceConfig, inputs []Topic) *wraperror.WrappedError {
	// Configure the Kafka admin client
	adminClient, err := kafka.NewAdminClient(
		&kafka.ConfigMap{"bootstrap.servers": conf.GetString("queue.host") + ":" + conf.GetString("queue.port")},
	)
	if err != nil {
		log.Fatalf("Failed to create admin client: %v", err)
	}

	var topics []kafka.TopicSpecification
	for _, topic := range inputs {
		// Specify the topic configuration
		topicConfig := kafka.TopicSpecification{
			Topic:             topic.TopicName,
			NumPartitions:     topic.NumberOfPartition,
			ReplicationFactor: topic.ReplicationFactor,
		}
		topics = append(topics, topicConfig)
	}

	// Create the topic
	results, err := adminClient.CreateTopics(context.TODO(), topics)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	// Check the result for success
	for _, result := range results {
		if result.Error.Error() != "" {
			log.Fatalf("Failed to create topic: %v", result.Error)
		}
		fmt.Printf("Topic %s created\n", result.Topic)
	}

	// Close the admin client
	adminClient.Close()
	return nil
}
