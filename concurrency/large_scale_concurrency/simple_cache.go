package large_scale_concurrency

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
			}
		}
	}
}

func (c *Cache) put(key []byte, value []byte) {
	segmentId, slotId, slotIndex, rbIndex := c.getId(key)
	slotIndex = (slotIndex + 1) % rds
	rbIndex = rbIndex + slotIndex
	c.locks[segmentId].Lock()
	c.segments[segmentId].slots[slotId].point = slotIndex
	c.segments[segmentId].rb[rbIndex].key = c.segments[segmentId].rb[rbIndex].key[:0]
	c.segments[segmentId].rb[rbIndex].value = c.segments[segmentId].rb[rbIndex].value[:0]
	c.segments[segmentId].rb[rbIndex].key = append(c.segments[segmentId].rb[rbIndex].key, key...)
	c.segments[segmentId].rb[rbIndex].value = append(c.segments[segmentId].rb[rbIndex].value, value...)
	c.locks[segmentId].Unlock()
}

func (c *Cache) get(key []byte) []byte {
	segmentId, slotId, slotIndex, rbIndex := c.getId(key)
	c.locks[segmentId].Lock()
	var value []byte
	for i := 0; i < rds; i++ {
		index := (slotIndex - i + rds) % rds
		if bytes.Equal(c.segments[segmentId].rb[rbIndex+index].key, key) {
			value = c.segments[segmentId].rb[rbIndex+index].value
			slotIndex = (slotIndex + 1) % rds
			rbIndex = rbIndex + slotIndex
			c.segments[segmentId].slots[slotId].point = slotIndex
			c.segments[segmentId].rb[rbIndex].key = c.segments[segmentId].rb[rbIndex].key[:0]
			c.segments[segmentId].rb[rbIndex].value = c.segments[segmentId].rb[rbIndex].value[:0]
			c.segments[segmentId].rb[rbIndex].key = append(c.segments[segmentId].rb[rbIndex].key, key...)
			c.segments[segmentId].rb[rbIndex].value = append(c.segments[segmentId].rb[rbIndex].value, value...)
			break
		}
	}
	c.locks[segmentId].Unlock()
	return value
}

func (c *Cache) getId(key []byte) (uint64, uint64, int, int) {
	hash := &maphash.Hash{}
	hash.SetSeed(c.seed)
	hash.Write(key)
	keyHash := hash.Sum64()
	segmentId := keyHash % segmentSize
	slotId := keyHash / segmentSize % slotSize
	slotIndex := c.segments[segmentId].slots[slotId].point
	rbIndex := int(slotId) * rds
	return segmentId, slotId, slotIndex, rbIndex
}

