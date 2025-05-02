package subpub

import (
	"context"
	"sync"
)

var (
	_ SubPub = (*PubSub)(nil)
)

type PubSub struct {
	mu         sync.RWMutex
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
				ps.mu.Lock()
				ps.Subscribed.safeDelete(se)
				ps.mu.Unlock()
				return
			default:
				continue
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

	ps.publishData(partitions, msg)

	return nil
}

func (ps *PubSub) publishData(partitions map[int]*subEntity, msg interface{}) {
	wg := sync.WaitGroup{}
	wg.Add(len(partitions))
	for i := range partitions {
		go func(id int) {
			defer wg.Done()
			partition := partitions[i]
			partition.queue <- msg
		}(i)
	}
	wg.Wait()
}

func (ps *PubSub) Close(ctx context.Context) error {
	ps.Subscribed.cleanup()

	if err := ctx.Err(); err != nil {
		return err
	}

	return nil
}

func NewPubSub() SubPub {
	s := &PubSub{
		Subscribed: newSubscribers(),
	}

	return s
}
