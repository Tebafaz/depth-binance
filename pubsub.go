package main

import (
	"sync"
)

type Pubsub struct {
	mu     sync.RWMutex
	subs   map[string][]chan interface{}
	closed bool
}

func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan interface{})
	return ps
}

func (ps *Pubsub) Subscribe(topic string) chan interface{} {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan interface{}, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *Pubsub) Publish(topic string, data interface{}) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		go func(ch chan interface{}) {
			ch <- data
		}(ch)
	}
}

func (ps *Pubsub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}

func (ps *Pubsub) HaveSubscriber(topic string) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.subs[topic]; ok {
		return true
	}
	return false
}

func (ps *Pubsub) ShowAllSubscribers() map[string][]chan interface{} {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	return ps.subs
}

func (ps *Pubsub) Unsubscribe(topic string, ch chan interface{}) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for index, value := range ps.subs[topic] {
		if ch == value {
			//close(ps.subs[topic][index])
			ps.subs[topic] = removeChannel(ps.subs[topic], index)
		}
	}
}

func removeChannel(slice []chan interface{}, index int) []chan interface{} {
	return append(slice[:index], slice[index+1:]...)
}
