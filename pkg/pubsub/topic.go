package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/mpsdantas/bottle/pkg/log"
)

type Topic interface {
	Publish(ctx context.Context, event string, value interface{})
}

type topic struct {
	tp *pubsub.Topic
}

func (t *topic) Publish(ctx context.Context, event string, value interface{}) {
	go func(c context.Context) {
		data, err := json.Marshal(value)
		if err != nil {
			log.Error(c, "could not marshal data",
				log.Err(err),
			)
			return
		}

		result := t.tp.Publish(c, &pubsub.Message{
			Data: data,
			Attributes: map[string]string{
				"x-event": event,
			},
		})

		id, err := result.Get(c)
		if err != nil {
			log.Error(c, "could not publish message data",
				log.Err(err),
			)
			return
		}

		log.Info(c, "message published successfully",
			log.String("id", id),
		)
	}(WithoutCancel(ctx))
}
