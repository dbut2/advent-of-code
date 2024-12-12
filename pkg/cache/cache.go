package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/fnv"
	"reflect"
	"sync"
)

type Cache struct {
	internal map[string]map[uint64]any
	mu       *sync.RWMutex
	stats    Stats
}

type Stats struct {
	Hit  int
	Miss int
}

func New() *Cache {
	return &Cache{
		internal: make(map[string]map[uint64]any),
		mu:       &sync.RWMutex{},
	}
}

func (c *Cache) Call(f any, v ...any) []any {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		panic(fmt.Sprintf("can only cache function calls, got %T", f))
	}

	if _, ok := c.internal[reflect.TypeOf(f).String()]; !ok {
		c.internal[reflect.TypeOf(f).String()] = make(map[uint64]any)
	}

	h := hash(v)

	if val, ok := c.internal[reflect.TypeOf(f).String()][h]; ok {
		c.stats.Hit++
		return val.([]any)
	}
	c.stats.Miss++

	val := reflect.ValueOf(f)

	var parameters []reflect.Value
	for _, p := range v {
		parameters = append(parameters, reflect.ValueOf(p))
	}

	ret := val.Call(parameters)

	var out []any
	for _, r := range ret {
		out = append(out, r.Interface())
	}

	c.internal[reflect.TypeOf(f).String()][h] = out
	return out
}

func (c *Cache) Stats() Stats {
	return c.stats
}

func hash(input interface{}) uint64 {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(input)
	if err != nil {
		panic(err.Error())
	}

	h := fnv.New64a()
	_, err = h.Write(buf.Bytes())
	if err != nil {
		panic(err.Error())
	}

	return h.Sum64()
}
