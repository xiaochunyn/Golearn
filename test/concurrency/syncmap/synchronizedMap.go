package main

import "sync"

type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// NewSynchronizedMap init a SynchronizedMap
func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

// Put puts a data to sm
func (sm SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

// Get gets a data from sm
func (sm SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	return sm.data[k]
}

// Delete deletes
func (sm SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	delete(sm.data, k)
}

// Each traverse map and pass k,v to callback function cb
// so that cb can operate the map
func (sm SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}
