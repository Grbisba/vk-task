package subpub

import (
	"context"
	"testing"

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
		sp := NewPubSub()
		assert.NotNil(t, sp)
	})
}

func TestPubSub_Subscribe(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewPubSub()

		sub, err := sp.Subscribe(subjectName, handlerFunc)
		assert.NoError(t, err)
		assert.NotPanics(t, func() {
			sub.Unsubscribe()
		})
	})
}

func TestPubSub_Publish(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewPubSub()

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
			sp := NewPubSub()

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
			sp := NewPubSub()

			err := sp.Publish(subjectName, message)
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errNoSubscriber)
			}
		})
	})
}

func TestPubSub_Close(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		sp := NewPubSub()

		ctx := context.Background()

		err := sp.Close(ctx)
		assert.NoError(t, err)
	})
	t.Run("negative:context is canceled", func(t *testing.T) {
		sp := NewPubSub()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := sp.Close(ctx)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, context.Canceled)
		}
	})
}
