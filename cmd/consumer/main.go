package main

import (
	"fmt"

	"github.com/gsouza97/go-expert-utils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Abre o Canal do RabbitMQ
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	// Cria um canal go para receber as mensagens
	msgs := make(chan amqp.Delivery)
	// Inicia o consumo das mensagens enviadas para o canal go em uma goroutine
	go rabbitmq.Consume(ch, msgs)
	// Loop infinito para receber as mensagens e printar no console
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		// Confirma o recebimento da mensagem para não colocar na fila novamente
		msg.Ack(false)
	}
}
