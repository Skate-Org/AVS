package mem

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/Skate-Org/AVS/lib/crypto/ecdsa"
	"github.com/Skate-Org/AVS/relayer/db/skateapp"
	"github.com/coocood/freecache"
	"github.com/pkg/errors"
)

var SEPARATOR = []byte("#")

type MemCache struct {
	*freecache.Cache
}

func NewCache(size int) *MemCache {
	return &MemCache{Cache: freecache.NewCache(size)}
}

type (
	Signature = skateapp.SignatureTuple
	Message   = skateapp.Message
)

func GenKey(entry Message) []byte {
	taskIdBytes := []byte(strconv.FormatUint(uint64(entry.TaskId), 10))
	chainTypeBytes := []byte(strconv.FormatUint(uint64(entry.ChainType), 10))
	chainIdBytes := []byte(strconv.FormatUint(uint64(entry.ChainId), 10))

	return ecdsa.Keccak256(taskIdBytes, chainTypeBytes, chainIdBytes)
}

func (cache *MemCache) CacheMessage(key []byte, entry Message) error {
	// NOTE: json encode for small size data (> gob), find more efficient way in futures
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	cacheKey := ecdsa.Keccak256([]byte("skateapp:task:"), key)
	return cache.Set(cacheKey, data, 0)
}

func (cache *MemCache) AppendSignature(key []byte, sig Signature) error {
	sigData, err := json.Marshal(sig)
	if err != nil {
		return err
	}
	cacheKey := ecdsa.Keccak256([]byte("skateapp:signed_task:"), key)
	existingData, err := cache.Get(cacheKey)
	if err == nil && len(existingData) > 0 {
		existingData = append(existingData, SEPARATOR...)
	}
	existingData = append(existingData, sigData...)
	return cache.Set([]byte(cacheKey), existingData, 0)
}

func (cache *MemCache) GetMessage(key []byte) (*Message, error) {
	cacheKey := ecdsa.Keccak256([]byte("skateapp:task:"), key)
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetMessage/Get")
	}
	if len(data) == 0 {
		return nil, nil
	}
	var result Message
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetMessage/UnMarshal")
	}

	return &result, nil
}

// GetSignatures retrieves and deserializes all signatures from the cache.
func (cache *MemCache) GetSignatures(key []byte) ([]Signature, error) {
	cacheKey := ecdsa.Keccak256([]byte("skateapp:signed_task:"), key)
	data, err := cache.Get([]byte(cacheKey))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	// Split the data using the separator
	parts := bytes.Split(data, SEPARATOR)
	signatures := make([]Signature, len(parts))

	for i, part := range parts {
		var sig Signature
		if err := json.Unmarshal(part, &sig); err != nil {
			return nil, err
		}
		signatures[i] = sig
	}

	return signatures, nil
}
