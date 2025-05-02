package subpub

import (
	"sync"
)

var _ Subscription = (*subEntity)(nil)

type subEntity struct {
	mh    MessageHandler
	name  string
	close chan struct{}
	queue chan interface{}
	id    int
}

func (s *subEntity) Unsubscribe() {
	s.close <- struct{}{}
}

func newSubEntity(subject string, mh MessageHandler) *subEntity {
	return &subEntity{
		mh:    mh,
		name:  subject,
		close: make(chan struct{}),
		queue: make(chan interface{}),
	}
}

type Subscribers struct {
	mu   sync.RWMutex
	subs map[string]map[int]*subEntity
}

func newSubscribers() *Subscribers {
	return &Subscribers{
		mu:   sync.RWMutex{},
		subs: make(map[string]map[int]*subEntity),
	}
}

func (s *Subscribers) add(entity *subEntity) {
	s.mu.Lock()
	defer s.mu.Unlock()

	partitions, ok := s.subs[entity.name]
	if !ok {
		partitions = make(map[int]*subEntity)
	}

	entity.id = len(partitions)
	partitions[entity.id] = entity
	s.subs[entity.name] = partitions
}

func (s *Subscribers) get(subject string) map[int]*subEntity {
	s.mu.RLock()
	se, ok := s.subs[subject]
	s.mu.RUnlock()

	if !ok {
		return nil
	}

	return se
}

func (s *Subscribers) safeDelete(se *subEntity) {
	s.mu.Lock()
	delete(s.subs[se.name], se.id)
	s.mu.Unlock()
}

func (s *Subscribers) cleanup() {
	s.mu.Lock()
	s.subs = make(map[string]map[int]*subEntity)
	s.mu.Unlock()
}
