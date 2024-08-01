package pokecash

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}
func TestGet(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Errorf("actual: %v, expected: 'val1'", actual)
	}
}
func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond*2)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped.", keyOne)
	}
}
