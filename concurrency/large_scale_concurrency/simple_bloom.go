package large_scale_concurrency

import (
	"hash/maphash"
	"math"
	"sync"
)

type Bloom struct {
	bits   []byte
	seeds  []maphash.Seed
	locker *sync.RWMutex
}

func NewBloom(n int) *Bloom {
	bits := make([]byte, math.MaxUint32)
	seeds := make([]maphash.Seed, n)
	locker := &sync.RWMutex{}
	for i := 0; i < n; i++ {
		seeds[i] = maphash.MakeSeed()
	}
	return &Bloom{bits: bits, seeds: seeds, locker: locker}
}

func (b *Bloom) Add(input string) error {
	for i := 0; i < len(b.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(b.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return err
		}
		key := hash.Sum64()
		index := key >> 32
		pos := key>>29 & 0x07
		b.locker.Lock()
		b.bits[index] |= 1 << pos
		b.locker.Unlock()
	}
	return nil
}

func (b *Bloom) IsExist(input string) (bool, error) {
	for i := 0; i < len(b.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(b.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return false, err
		}
		key := hash.Sum64()
		index := key >> 32
		pos := key>>29 & 0x07
		b.locker.RLock()
		existed := b.bits[index]&(1<<pos) != 0
		b.locker.RUnlock()
		return existed, nil
	}
	return true, nil
}
