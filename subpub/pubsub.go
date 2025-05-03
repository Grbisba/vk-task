package subpub

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	_ SubPub = (*PubSub)(nil)
)

type PubSub struct {
	mu         sync.Mutex
	Subscribed *Subscribers
}

// MessageHandler is a callback function that processes messages delivered to subscribers.
type MessageHandler func(msg interface{})

// Subscribe should subscribe to name and consume messages.
func (ps *PubSub) Subscribe(subject string, mh MessageHandler) (Subscription, error) {
	se := newSubEntity(subject, mh)

	ps.mu.Lock()
	ps.Subscribed.add(se)
	ps.mu.Unlock()

	go func() {
		for {
			select {
			case v := <-se.queue:
				se.mh(v)
			case <-se.close:
				for v := range se.queue {
					se.mh(v)
				}
				ps.Subscribed.safeDelete(se)
				return
			}
		}
	}()

	return se, nil
}

// Publish should send messages to the name.
func (ps *PubSub) Publish(subject string, msg interface{}) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	partitions := ps.Subscribed.get(subject)
	if partitions == nil || len(partitions) == 0 {
		return errNoSubscriber
	}

	err := ps.publishData(partitions, msg)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PubSub) publishData(partitions map[int]*subEntity, msg interface{}) error {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(partitions))
	for i := range partitions {
		go func(id int) {
			defer wg.Done()
			partition := partitions[i]
			if partition.closed {
				err = errNoSubscriber
				return
			}
			select {
			case partition.queue <- msg:
			case <-time.After(time.Second):
				fmt.Println("TIMEOUT")
			}
		}(i)
	}
	wg.Wait()

	return err
}

func (ps *PubSub) Close(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return ctx.Err()
	}

	ps.mu.Lock()
	for _, partition := range ps.Subscribed.subs {
		for _, se := range partition {
			se.Unsubscribe()
		}
	}
	ps.mu.Unlock()

	return nil
}

func NewSubPub() SubPub {
	s := &PubSub{
		Subscribed: newSubscribers(),
	}

	return s
}
