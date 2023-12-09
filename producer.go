package main

import (
	"github.com/IBM/sarama"
	"log"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil) //the address and the error
	if err != nil {
		log.Fatalf("Failed to create a kafka producer %s", err)

	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {

		}
	}(producer)

	message := &sarama.ProducerMessage{Topic: "My-Message",
		Value: sarama.StringEncoder("Hello from Kafka"),
	}

	_, _, err = producer.SendMessage(message) //partition offset and error
	if err != nil {
		log.Fatalf("Error sending the message %s", err)

	}
	log.Println("Message sent successfully! ")

}
