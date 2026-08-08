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

func InitDataStream(ctx context.Context, js jetstream.JetStream) error {
	streams := []DataStreamConfig{
		{Name: ExecDS.Name, Subject: ExecDS.Subject},
		{Name: ExecveDS.Name, Subject: ExecveDS.Subject},
		{Name: DbusDS.Name, Subject: DbusDS.Subject},
		{Name: Connect4DS.Name, Subject: Connect4DS.Subject},
		{Name: Bind4DS.Name, Subject: Bind4DS.Subject},
		{Name: ISSSDS.Name, Subject: ISSSDS.Subject},
		{Name: FanotifyDS.Name, Subject: FanotifyDS.Subject},
	}

	for _, s := range streams {
		_, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
			Name:     s.Name,
			Subjects: []string{s.Subject},
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error creating stream: %w", err)
		}
	}

	return nil
}
