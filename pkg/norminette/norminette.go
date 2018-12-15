package norminette

import "github.com/streadway/amqp"

type Norminette struct {
	Connection *amqp.Connection
}

func New(uri string) (*Norminette, error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return &Norminette{Connection: connection}, nil
}
