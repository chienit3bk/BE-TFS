package producer

import (
	"context"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type Producer struct {
	exchange     string
	exchangeType string
	routingKey   string
	channel      *amqp.Channel
	ctx          context.Context
	wg           *sync.WaitGroup
}

func CreateNewProducer(exchange, exchangeType, routingKey string, channel *amqp.Channel, ctx context.Context,
	wg *sync.WaitGroup) *Producer {
	return &Producer{
		exchange:     exchange,
		exchangeType: exchangeType,
		routingKey:   routingKey,
		channel:      channel,
		ctx:          ctx,
		wg:           wg,
	}
}

func (p *Producer) Send(msg chan string) {
	if p.exchange == "" || p.exchangeType == "" || p.channel == nil {
		fmt.Println("This Producer has a faulty configuration")
		return
	}
	p.declare()
	for {
		select {
		case m := <-msg:
			fmt.Printf("Sending to %v : %v\n", p.exchange, m)
			err := p.publish(m)
			if err != nil {
				fmt.Println("Publish msg error: ", err)
			}
		case <-p.ctx.Done():
			fmt.Printf("Publish to exhange %v is stopped \n", p.exchange)
			p.wg.Done()
			return
		}
	}

}
func (p *Producer) publish(msg string) error {
	err := p.channel.Publish(
		p.exchange,
		p.routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plane",
			ContentEncoding: "",
			Body:            []byte(msg),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func (p *Producer) Close() error {
	return p.channel.Close()
}

//hàm khai báo exchange: hàm này không cần thiết: có hay không thì nó vẫn chạy
func (p *Producer) declare() error {
	err := p.channel.ExchangeDeclare(
		p.exchange,     // name
		p.exchangeType, //type
		true,           //durable
		false,          //autoDelete: delete when complete
		false,          //internal
		false,          //noWait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}
	return nil
}
