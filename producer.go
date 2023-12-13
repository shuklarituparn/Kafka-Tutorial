package main

import (
	"bufio"
	"github.com/IBM/sarama"
	"log"
	"os"
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

	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for scanner.Scan() {
		text = scanner.Text()
		message := &sarama.ProducerMessage{Topic: "My-Message",
			Value: sarama.StringEncoder(text),
		}

		_, _, err = producer.SendMessage(message) //partition offset and error
		if err != nil {
			log.Fatalf("Error sending the message %s", err)

		}
		log.Println("Message sent successfully! ")
	}

}
