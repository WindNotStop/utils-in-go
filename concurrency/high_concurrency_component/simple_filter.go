package high_concurrency_component

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

func (f *Filter) Add(input string) error {
	for i := 0; i < len(f.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(f.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return err
		}
		key := hash.Sum64()
		value, loaded := f.m.LoadOrStore(key, 1)
		if loaded {
			f.m.Store(key, value.(int)+1)
		}
	}
	return nil
}

func (f *Filter) IsExist(input string) (bool, error) {
	for i := 0; i < len(f.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(f.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return false, err
		}
		key := hash.Sum64()
		if v, existed := f.m.Load(key); !existed {
			if v == 0 {
				f.m.Delete(key)
			}
			return false, nil
		}
	}
	return true, nil
}

func (f *Filter) Remove(input string) error {
	for i := 0; i < len(f.seeds); i++ {
		hash := &maphash.Hash{}
		hash.SetSeed(f.seeds[i])
		_, err := hash.WriteString(input)
		if err != nil {
			return err
		}
		key := hash.Sum64()
		value, ok := f.m.Load(key)
		if !ok {
			return errors.New("not existed")
		}
		if value == 1 {
			f.m.Delete(key)
		}else {
			f.m.Store(key, value.(int)-1)
		}
	}
	return nil
}

func (f *Filter) String() string {
	sb := &strings.Builder{}
	f.m.Range(func(key, value interface{}) bool {
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
