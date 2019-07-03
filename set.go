package ryutil

import (
	"sync"
)

type ConcurrentSliceSet struct {
	m map[string]bool
	sync.RWMutex
}

func NewConcurrentSlicePair() *ConcurrentSliceSet {
	return &ConcurrentSliceSet{
		m: map[string]bool{},
	}
}

func (s *ConcurrentSliceSet) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *ConcurrentSliceSet) Remove(item string) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *ConcurrentSliceSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *ConcurrentSliceSet) Len() int {
	return len(s.List())
}

func (s *ConcurrentSliceSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *ConcurrentSliceSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *ConcurrentSliceSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (cs *ConcurrentSliceSet) Iter() <-chan string {
	c := make(chan string)

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, _ := range cs.m {
			c <- index
		}
		close(c)
	}
	go f()

	return c
}




func (cs *ConcurrentSliceSet) IterFn(funch func(k string))  {
	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, _ := range cs.m {
			funch(index)
		}
	}
	go f()
}


