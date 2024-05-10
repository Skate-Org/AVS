package mem

import (
	// "math/big"
	"math/big"
	"reflect"
	"testing"
)

// TestCacheOperator tests caching an entry.
func TestCacheOperator(t *testing.T) {
	cache := NewCache(2 * 1024 * 1024)
	entry := Operator{
		Address: "0x123",
		// Strategy: "USDC",
		// Amount:   "123333333333333333333",
	}
	if err := cache.CacheOperator(entry); err != nil {
		t.Errorf("CacheMessage() error = %v", err)
	}

	cachedEntry, err := cache.GetOperator(entry.Address)
	if err != nil {
		t.Errorf("Failed to get data from cache: %v", err)
		return
	}

	if !reflect.DeepEqual(*cachedEntry, entry) {
		t.Errorf("Cached entry %v does not match original entry %v", cachedEntry, entry)
	}

	evicted := cache.EvictOperator(entry.Address)
	cachedEntry, _ = cache.GetOperator(entry.Address)
	if cachedEntry != nil {
		t.Errorf("Cached entry %v still persist in cache", cachedEntry)
	}
	if !evicted {
		t.Errorf("Eviction status error, should be true")
	}
}

func TestCacheOperatorCount(t *testing.T) {
	cache := NewCache(2 * 1024 * 1024)

	value := uint32(100000)
	if err := cache.CacheOperatorCount(value); err != nil {
		t.Errorf("CacheMessage() error = %v", err)
	}

	cachedEntry, err := cache.GetOperatorCount()
	// t.Logf("Result from cache: %v", *cachedEntry)
	if err != nil {
		t.Errorf("Failed to get data from cache: %v", err)
		return
	}

	if !reflect.DeepEqual(*cachedEntry, value) {
		t.Errorf("Cached entry %v does not match original entry %v", cachedEntry, value)
	}
}

func TestCacheStake(t *testing.T) {
	cache := NewCache(2 * 1024 * 1024)

	strategy := "stETH"
	value := big.NewInt(100000)
	if err := cache.CacheStake(strategy, *value); err != nil {
		t.Errorf("CacheMessage() error = %v", err)
	}

	cachedEntry, err := cache.GetStake(strategy)
	// t.Logf("Result from cache: %v", *cachedEntry)
	if err != nil {
		t.Errorf("Failed to get data from cache: %v", err)
		return
	}

	if !reflect.DeepEqual(*cachedEntry, *value) {
		t.Errorf("Cached entry %v does not match original entry %v", cachedEntry, value)
	}
}
