package mem

import (
	"reflect"
	"testing"

	pb "github.com/Skate-Org/AVS/api/pb/relayer"
)

// TestCacheMessage tests caching an entry.
func TestCacheMessage(t *testing.T) {
	cache := NewCache(2 * 1024 * 1024)
	entry := Message{
		TaskId:    1,
		Initiator: "123 github.com/Skate-Org/AVS",
		Message:   "Test Entry",
		ChainType: pb.ChainType_EVM,
		ChainId:   1,
	}
	key := GenKey(entry)

	if err := cache.CacheMessage(key, entry); err != nil {
		t.Errorf("CacheMessage() error = %v", err)
	}

	cachedEntry, err := cache.GetMessage(key)
	if err != nil {
		t.Errorf("Failed to get data from cache: %v", err)
		return
	}

	if !reflect.DeepEqual(*cachedEntry, entry) {
		t.Errorf("Cached entry %v does not match original entry %v", cachedEntry, entry)
	}
}

// TestAppendSignature tests appending a signature to the cache.
func TestAppendSignature(t *testing.T) {
	cache := NewCache(2 * 1024 * 1024)
	// Clear previous data
	refEntry := Message{
		TaskId:    1,
		Initiator: "123 github.com/Skate-Org/AVS",
		Message:   "Test Entry",
		ChainType: pb.ChainType_EVM,
		ChainId:   1,
	}
	key := GenKey(refEntry)

	signatures := []Signature{
		{Operator: "0x1_github.com/Skate-Org/AVS", Signature: [65]byte{1}},
		{Operator: "0x2_github.com/Skate-Org/AVS", Signature: [65]byte{2}},
		{Operator: "0x3_github.com/Skate-Org/AVS", Signature: [65]byte{3}},
	}
	for _, sig := range signatures {
		if err := cache.AppendSignature(key, sig); err != nil {
			t.Errorf("Failed to append signature: %v", err)
			return
		}
	}

	cachedSignatures, err := cache.GetSignatures(key)
	if err != nil {
		t.Errorf("GetSignatures() error = %v", err)
		return
	}

	if len(cachedSignatures) != 3 || !reflect.DeepEqual(signatures, cachedSignatures) {
		t.Errorf("Appended signature %v does not match retrieved signature %v", signatures, cachedSignatures)
	}
}
