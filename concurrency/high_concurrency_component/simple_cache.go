package high_concurrency_component

import (
	"bytes"
	"hash/maphash"
	"sync"
)

const (
	segmentSize = 256
	slotSize    = 128
	rbSize      = 8192
	rds         = rbSize / slotSize
	keyLen      = 128
	valueLen    = 128
)

type Cache struct {
	locks    []sync.Mutex
	segments []segment
	seed     maphash.Seed
}

type segment struct {
	rb    []entity
	slots []slot
}

type slot struct {
	offset int
	point  int
}

type entity struct {
	time  byte
	key   []byte
	value []byte
}

func (c *Cache) Init() {
	c.locks = make([]sync.Mutex, segmentSize)
	c.segments = make([]segment, segmentSize)
	c.seed = maphash.MakeSeed()
	for i := 0; i < segmentSize; i++ {
		c.locks[i] = sync.Mutex{}
		c.segments[i].rb = make([]entity, rbSize)
		c.segments[i].slots = make([]slot, slotSize)
		for j := 0; j < slotSize; j++ {
			c.segments[i].slots[j].offset = j * rds
			c.segments[i].slots[j].point = 0
			for k := 0; k < rds; k++ {
				index := j*rds + k
				c.segments[i].rb[index].key = make([]byte, 0, keyLen)
				c.segments[i].rb[index].value = make([]byte, 0, valueLen)
				c.segments[i].rb[index].time = 0
			}
		}
	}
}

func (c *Cache) Put(key []byte, value []byte) {
	segmentId, slotId := c.getId(key)
	c.locks[segmentId].Lock()
	slotIndex := c.segments[segmentId].slots[slotId].point
	rbIndex := int(slotId) * rds
Loop:
	for {
		for i := 1; i <= rds; i++ {
			index := (slotIndex+i+rds)%rds + rbIndex
			if c.segments[segmentId].rb[index].time <= 0 {
				c.segments[segmentId].slots[slotId].point = index - rbIndex
				c.segments[segmentId].rb[index].key = c.segments[segmentId].rb[index].key[:0]
				c.segments[segmentId].rb[index].value = c.segments[segmentId].rb[index].value[:0]
				c.segments[segmentId].rb[index].key = append(c.segments[segmentId].rb[index].key, key...)
				c.segments[segmentId].rb[index].value = append(c.segments[segmentId].rb[index].value, value...)
				c.segments[segmentId].rb[index].time = 99
				break Loop
			} else {
				c.segments[segmentId].rb[index].time--
			}
		}
	}
	c.locks[segmentId].Unlock()
}

func (c *Cache) Get(key []byte) []byte {
	segmentId, slotId := c.getId(key)
	c.locks[segmentId].Lock()
	slotIndex := c.segments[segmentId].slots[slotId].point
	rbIndex := int(slotId) * rds
	var value []byte
	for i := 0; i < rds; i++ {
		index := (slotIndex+i+rds)%rds + rbIndex
		if bytes.Equal(c.segments[segmentId].rb[index].key, key) {
			value = c.segments[segmentId].rb[index].value
			if c.segments[segmentId].rb[index].time < 100 {
				c.segments[segmentId].rb[index].time += 9
			}
			break
		} else {
			c.segments[segmentId].rb[index].time--
		}
	}
	c.locks[segmentId].Unlock()
	return value
}

func (c *Cache) getId(key []byte) (uint64, uint64) {
	hash := &maphash.Hash{}
	hash.SetSeed(c.seed)
	hash.Write(key)
	keyHash := hash.Sum64()
	segmentId := keyHash % segmentSize
	slotId := keyHash / segmentSize % slotSize
	return segmentId, slotId
}
