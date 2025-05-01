package pubsub

import (
	"sync"
)

type Subjects struct {
	mu   sync.RWMutex
	subs map[string]int
}

func NewSubjects() *Subjects {
	return &Subjects{
		mu:   sync.RWMutex{},
		subs: make(map[string]int),
	}
}

func (s *Subjects) AddSubject(subject string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subs[subject] = s.subs[subject] + 1
}

func (s *Subjects) RemoveSubject(subject string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subs[subject] = s.subs[subject] - 1
}

func (s *Subjects) Count(subject string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.subs[subject]
}
