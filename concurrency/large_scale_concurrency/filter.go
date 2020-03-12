package large_scale_concurrency

import (
	"errors"
	"hash/maphash"
	"strconv"
	"strings"
	"sync"
)

type Filter struct {
	m     *sync.Map
	seeds []maphash.Seed
}

func NewFilter(n int) *Filter {
	m := &sync.Map{}
	seeds := make([]maphash.Seed, n)
	for i := 0; i < n; i++ {
		seeds[i] = maphash.MakeSeed()
	}
	return &Filter{m: m, seeds: seeds}
}

func (b *Filter) Add(input string) error {
	for i := 0; i < len(b.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(b.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return err
		}
		key := hash.Sum64()
		value, loaded := b.m.LoadOrStore(key, 1)
		if loaded {
			b.m.Store(key, value.(int)+1)
		}
	}
	return nil
}

func (b *Filter) IsExist(input string) (bool, error) {
	ok := true
	for i := 0; i < len(b.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(b.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return false, err
		}
		key := hash.Sum64()
		if v, existed := b.m.Load(key); !existed {
			ok = false
			if v == 0 {
				b.m.Delete(key)
			}
		}
	}
	return ok, nil
}

func (b *Filter) Remove(input string) error {
	for i := 0; i < len(b.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(b.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return err
		}
		key := hash.Sum64()
		value, ok := b.m.Load(key)
		if !ok {
			return errors.New("not existed")
		}
		b.m.Store(key, value.(int)-1)
		if value == 1 {
			b.m.Delete(key)
		}
	}
	return nil
}

func (b *Filter) String() string {
	sb := &strings.Builder{}
	b.m.Range(func(key, value interface{}) bool {
		sb.WriteString("key:")
		sb.WriteString(strconv.FormatUint(key.(uint64), 10))
		sb.WriteString(" , ")
		sb.WriteString("value:")
		sb.WriteString(strconv.Itoa(value.(int)))
		sb.WriteByte(10)
		return true
	})
	return sb.String()
}
