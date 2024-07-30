package pubsub

import (
	"context"
	"encoding/json"
	"time"

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

		for i := 0; i < 100; i++ {
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

				time.Sleep(1 * time.Second)
				continue
			}

			log.Info(c, "message published successfully",
				log.String("id", id),
			)

			break
		}
	}(WithoutCancel(ctx))
}
