package common

import (
	"context"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

func NewDefaultSaramaProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Compression = sarama.CompressionSnappy

	return config
}

func NewDefaultSaramaConsumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	return config
}

type MessageHandler interface {
	ConsumeMessage(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage)
}

type Consumer struct {
	MessageHandler
	Ready chan bool
}

func NewConsumer(messageHandler MessageHandler) *Consumer {
	return &Consumer{
		Ready:          make(chan bool),
		MessageHandler: messageHandler,
	}
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		consumer.MessageHandler.ConsumeMessage(session, message)
	}

	return nil
}

func StartConsumer(ctx context.Context, client sarama.ConsumerGroup, topics string, consumer *Consumer) {
	go func() {
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
			consumer.Ready = make(chan bool)
		}
	}()
}
