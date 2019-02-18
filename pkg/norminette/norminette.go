package norminette

import "github.com/streadway/amqp"

type Norminette struct {
	Connection    *amqp.Connection
	DisabledRules []string
}

func New(uri string, disabledRules []string) (*Norminette, error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return &Norminette{
		Connection:    connection,
		DisabledRules: disabledRules,
	}, nil
}
