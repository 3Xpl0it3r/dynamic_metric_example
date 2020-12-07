package crcache

import "sync"

type CrCache struct {
	Data map[string]int32
	mux sync.Mutex
}


func NewCrCache()*CrCache{
	return &CrCache{
		Data: make(map[string]int32),
		mux:  sync.Mutex{},
	}
}

func(c *CrCache)Add(key string, value int32){
	c.mux.Lock()
	_, ok := c.Data[key]
	if !ok {
		c.Data[key] = value
	} else {
		c.Data[key] = c.Data[key] + 1
	}
	c.mux.Unlock()
}

func(c *CrCache)Update(key string, value int32){
	c.mux.Lock()
	c.Data[key] = value
	c.mux.Unlock()
}

func(c *CrCache)Remove(key string){
	c.mux.Lock()
	_, ok := c.Data[key]
	if ok {
		delete(c.Data, key)
	}
	c.mux.Unlock()
}

func(c *CrCache)Get(key string)(int32,bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.Data[key]

	if !ok {
		return 0, false
	}
	return c.Data[key], true
}
