package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Função que cria uma conexão com o RabbitMQ,
// criando também um canal para trabalhar com essa conexão
func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, err
}

func Consume(ch *amqp.Channel, msgsOut chan<- amqp.Delivery) error {
	// A função consume fica fazendo um pooling na fila do RabbitMQ pegando as informações
	msgs, err := ch.Consume("minha-fila", "go-consumer", false, false, false, false, nil)
	if err != nil {
		return err
	}
	// Começar a consumir as mensagens e jogar para o canal go
	for msg := range msgs {
		msgsOut <- msg
	}
	return nil
}
