package rabbitMQ

import (
	"context"
	"fmt"
	"project/packages/rabbitMQ/consumer"
	"project/packages/rabbitMQ/producer"
	"project/packages/rabbitMQ/rabbitmq"
	"sync"

)

func QuickCreateNewPairProducerAndConsumer(exchangeName, queueName string, ctx context.Context, wg *sync.WaitGroup) (*producer.Producer, *consumer.Consumer, error) {
	//trick configuration :v
	routingKey := "abc"
	bindingKey := routingKey
	exchangeType := "direct"
	//create RMQ connection
	rmq := rabbitmq.CreateNewRMQ(URI)
	if rmq == nil {
		return &producer.Producer{}, &consumer.Consumer{}, fmt.Errorf("initialization rmq failed")
	}

	// create 1 channel for producer
	pCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return &producer.Producer{}, &consumer.Consumer{}, err
	}
	// create 1 channel for consumer
	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return &producer.Producer{}, &consumer.Consumer{}, err
	}

	return producer.CreateNewProducer(exchangeName, exchangeType, routingKey, pCh, ctx, wg),
		consumer.CreateNewConsumer(exchangeName, exchangeType, bindingKey, queueName, cCh, ctx, wg),
		nil
}
