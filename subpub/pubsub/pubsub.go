package pubsub

import (
	"context"

	"github.com/vk-task/subpub"
)

var (
	_ subpub.SubPub       = (*S)(nil)
	_ subpub.Subscription = (*sub)(nil)
)

type sub struct {
	subjects Subjects
}

func (s *sub) Unsubscribe() {}

type S struct {
	SubjectName string
	// Subjects describes a subjects and provided consumers.
	Subjects Subjects
	Queue    chan interface{}
}

// Subscribe should subscribe to subject and consume messages.
func (s *S) Subscribe(subject string, cb subpub.MessageHandler) (subpub.Subscription, error) {
	return nil, nil
}

// Publish should send messages to the subject.
func (s *S) Publish(subject string, msg interface{}) error {
	return nil
}

func (s *S) Close(ctx context.Context) error {
	return nil
}

func NewSubPub() subpub.SubPub {
	s := &S{}

	return s
}
