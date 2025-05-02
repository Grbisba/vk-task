package subpub

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribers_addSubscription(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		s := newSubscribers()
		subject := "123"

		se := newSubEntity(subject, func(msg interface{}) {
			fmt.Printf("%s recieved message: %s\n", subject, msg)
		})

		s.add(se)

		assert.Equal(t, 1, len(s.subs[subject]))
		assert.Equal(t, s.get("123"), map[int]*subEntity{0: se})

		s.add(se)

		assert.Equal(t, 2, len(s.subs[subject]))
		assert.Equal(t, s.get("123"), map[int]*subEntity{0: se, 1: se})
	})
}

func TestSubscribers_getSubscription(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		s := newSubscribers()
		subject := "123"

		se := newSubEntity(subject, func(msg interface{}) {
			fmt.Printf("%s recieved message: %s\n", subject, msg)
		})

		s.add(se)

		assert.Equal(t, 1, len(s.subs[subject]))
		assert.Equal(t, s.get("123"), map[int]*subEntity{0: se})

		s.add(se)

		assert.Equal(t, 2, len(s.subs[subject]))
		assert.Equal(t, s.get("123"), map[int]*subEntity{0: se, 1: se})

		subs := s.get("123")
		assert.Equal(t, 2, len(subs))
		assert.Equal(t, subs[0], se)
		assert.Equal(t, subs[1], se)
	})
}
