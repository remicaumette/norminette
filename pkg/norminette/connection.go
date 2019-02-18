package norminette

import (
	"encoding/json"

	"github.com/remicaumette/norminette/pkg/protocol"
	"github.com/streadway/amqp"
)

func (norminette *Norminette) Publish(channel *amqp.Channel, queue *amqp.Queue, v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return channel.Publish("", "norminette", false, false, amqp.Publishing{
		ReplyTo: queue.Name,
		Body:    bytes,
	})
}

func (norminette *Norminette) InitConnection() (*amqp.Channel, *amqp.Queue, error) {
	channel, err := norminette.Connection.Channel()
	if err != nil {
		return nil, nil, err
	}
	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		return nil, nil, err
	}
	if err = channel.Qos(1, 0, false); err != nil {
		return nil, nil, err
	}
	return channel, &queue, nil
}

func (norminette *Norminette) Version() (*protocol.VersionResponse, error) {
	channel, queue, err := norminette.InitConnection()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err = norminette.Publish(channel, queue, protocol.NewVersionRequest()); err != nil {
		return nil, err
	}

	response := &protocol.VersionResponse{}
	message := <-messages
	if err = json.Unmarshal(message.Body, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (norminette *Norminette) CheckFiles(files ...string) ([]protocol.CheckFileResponse, error) {
	count := 0
	result := make([]protocol.CheckFileResponse, 0)
	channel, queue, err := norminette.InitConnection()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		request, err := protocol.NewCheckFileRequest(file, norminette.DisabledRules)
		if err != nil {
			return nil, err
		}
		if err = norminette.Publish(channel, queue, request); err != nil {
			return nil, err
		}
	}

	for message := range messages {
		response := protocol.CheckFileResponse{}
		if err = json.Unmarshal(message.Body, &response); err != nil {
			return nil, err
		}
		result = append(result, response)
		count++
		if count == len(files) {
			break
		}
	}
	return result, nil
}
