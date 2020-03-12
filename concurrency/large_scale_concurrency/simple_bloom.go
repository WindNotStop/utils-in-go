package large_scale_concurrency

import (
	"hash/maphash"
	"sync"
)

type Bloom struct {
	size   uint
	bits   []byte
	seeds  []maphash.Seed
	locker *sync.RWMutex
}

//1024大小的bloom的size输入为10，n为hash函数数量
func NewBloom(size uint, n int) *Bloom {
	bits := make([]byte, 1<<size-3)
	seeds := make([]maphash.Seed, n)
	locker := &sync.RWMutex{}
	for i := 0; i < n; i++ {
		seeds[i] = maphash.MakeSeed()
	}
	return &Bloom{size: size - 3, bits: bits, seeds: seeds, locker: locker}
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
		index := key >> uint(61-b.size)
		pos := key >> uint(64-b.size) & 0x07
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
		index := key >> uint(61-b.size)
		pos := key >> uint(64-b.size) & 0x07
		b.locker.RLock()
		existed := b.bits[index]&(1<<pos) != 0
		b.locker.RUnlock()
		return existed, nil
	}
	return true, nil
}
