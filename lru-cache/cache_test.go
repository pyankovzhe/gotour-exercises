package lrucache

import "testing"

func TestCacheSet(t *testing.T) {
	lru := NewCache(3)

	values := []struct {
		key  Key
		data interface{}
	}{
		{key: "foo", data: 1},
		{key: "bar", data: 2},
		{key: "foobar", data: 3},
	}

	t.Run("Set/Get logic", func(t *testing.T) {
		for _, value := range values {
			present := lru.Set(value.key, value.data)
			if present {
				t.Fatalf("Expected %v, got %v", false, present)
			}
		}

		for _, value := range values {
			cachedVal, cached := lru.Get(value.key)
			if !cached {
				t.Fatalf("Expected %v, got %v", value.data, cachedVal)
			}
		}
	})
}
