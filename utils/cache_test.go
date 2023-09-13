package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, err := cache.Get(c.key)
			if err != nil {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = time.Duration(1) * time.Second
	const waitTime = baseTime + 1*time.Second
	cache := NewCache(baseTime)
	key := "https://example.com"
	cache.Add(key, []byte("testdata"))

	val, err := cache.Get(key)
	if err != nil {
		t.Errorf("expected to find key")
		return
	}
	fmt.Println("Found", val, "For key", key)
	time.Sleep(waitTime)

	val2, err := cache.Get(key)
	if err == nil {
		fmt.Println("Found", val2, "For key", key)
		t.Errorf("expected to not find key")
		return
	}
}
