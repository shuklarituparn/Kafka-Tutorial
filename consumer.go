package main

import (
	"github.com/IBM/sarama"
	"log"
	"sync"
)

func main() {

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating a consumer %s", err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalf("Error closing the consumer %s", err)
		}
	}(consumer)

	//Subscribing to the topic
	topic := "My-Message"
	paritions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get the partition for the topic %s: %s", topic, err)
	}
	var wg sync.WaitGroup
	wg.Add(len(paritions))

	for _, partition := range paritions {
		go func(partition int32) {
			defer wg.Done()
			partitionConsume, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Fatalf("Failed to create consumer partition %s", err)

			}
			defer func(partitionConsume sarama.PartitionConsumer) {
				err := partitionConsume.Close()
				if err != nil {

				}
			}(partitionConsume)
			for message := range partitionConsume.Messages() {
				log.Printf("Recieved the message %s", string(message.Value))

			}
		}(partition)

	}
	wg.Wait()
}
