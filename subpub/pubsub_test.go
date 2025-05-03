package subpub

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	subjectName = "testSubject"
	message     = "testMessage"
)

var (
	handlerFunc = func(msg interface{}) {}
)

func TestPubSub_NewPubSub(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewSubPub()
		assert.NotNil(t, sp)
	})
}

func TestPubSub_Subscribe(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewSubPub()

		sub, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)
		assert.NotPanics(t, func() {
			sub.Unsubscribe()
		})
	})
}

func TestPubSub_Publish(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewSubPub()

		sub, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)

		assert.NotPanics(t, func() {
			sub.Unsubscribe()
		})
	})
	t.Run("negative:errNoSubscribers", func(t *testing.T) {
		t.Run("case:1", func(t *testing.T) {
			sp := NewSubPub()

			sub, err := sp.Subscribe(subjectName, handlerFunc)
			assert.NoError(t, err)
			assert.NotPanics(t, func() {
				sub.Unsubscribe()
			})

			err = sp.Publish(subjectName, message)
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errNoSubscriber)
			}
		})
		t.Run("case:2", func(t *testing.T) {
			sp := NewSubPub()

			err := sp.Publish(subjectName, message)
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errNoSubscriber)
			}
		})
	})
}

func TestPubSub_Close(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := &PubSub{
			Subscribed: newSubscribers(),
		}

		_, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)

		ctx := context.Background()

		err = sp.Close(ctx)
		assert.NoError(t, err)

		// should wait while subscribe collect data and close chan
		time.Sleep(100 * time.Millisecond)

		assert.Len(t, sp.Subscribed.subs[subjectName], 0)
	})
	t.Run("negative:context is canceled", func(t *testing.T) {
		sp := &PubSub{
			Subscribed: newSubscribers(),
		}

		_, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)

		ctx := context.Background()
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)

		cancel()

		err = sp.Close(ctx)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, context.Canceled)
		}

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)
	})
	t.Run("negative:context deadline exceed", func(t *testing.T) {
		sp := &PubSub{
			Subscribed: newSubscribers(),
		}

		_, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)

		ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)

		err = sp.Close(ctx)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		}

		err = sp.Publish(subjectName, message)
		assert.NoError(t, err)
	})
}
