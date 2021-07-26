package realtime

const (
	exchangeName = "productSender"
	exchangeType = "direct"
	queueName    = "productReceiver"
	URI          = "amqp://tfs:tfs-ocg@174.138.40.239:5672/#/"
	routingKey   = "awm"
	bindingKey   = "awm"
)

var (
	productCounterOfSender   = 0
	productCounterOfReceiver = 0
)
