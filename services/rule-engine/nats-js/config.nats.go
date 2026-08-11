package natsjs

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var NatsInstance *nats.Conn
var NatsJSInstance jetstream.JetStream

func Get() (*nats.Conn, jetstream.JetStream, error) {
	// nats
	if NatsInstance == nil {
		return nil, nil, fmt.Errorf("[ERROR] NATS connection is not initialized yet")
	}

	// js
	if NatsJSInstance == nil {
		return NatsInstance, nil, fmt.Errorf("[ERROR] JetStream connection is not initialized yet")
	}

	return NatsInstance, NatsJSInstance, nil
}

func Connect(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Connecting with NATS has been failed: %w", err)
	}

	NatsInstance = nc
	return nc, nil
}

func JSConnect(nc *nats.Conn) (jetstream.JetStream, error) {
	jc, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Connecting with JetStream has been failed: %w", err)
	}

	NatsJSInstance = jc
	return jc, nil
}

func InitConsumer(ctx context.Context, js jetstream.JetStream) error {
	streams := []DataStreamConfig{
		{Name: ExecDS.Name, Durable: ExecDS.Durable},
		{Name: ExecveDS.Name, Durable: ExecveDS.Durable},
		{Name: DbusDS.Name, Durable: DbusDS.Durable},
		{Name: Connect4DS.Name, Durable: Connect4DS.Durable},
		{Name: Bind4DS.Name, Durable: Bind4DS.Durable},
		{Name: ISSSDS.Name, Durable: ISSSDS.Durable},
		{Name: FanotifyDS.Name, Durable: FanotifyDS.Durable},
	}

	for _, s := range streams {
		_, err := js.CreateOrUpdateConsumer(ctx, s.Name, jetstream.ConsumerConfig{
			Durable:       s.Durable,
			DeliverPolicy: jetstream.DeliverAllPolicy,
			AckPolicy:     jetstream.AckExplicitPolicy,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error creating consumer: %w", err)
		}
	}

	return nil
}
