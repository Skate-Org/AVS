package mem

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/big"

	"github.com/coocood/freecache"
	"github.com/pkg/errors"
	libHash "github.com/Skate-Org/AVS/lib/crypto/hash"
)

type MemCache struct {
	*freecache.Cache
}

func NewCache(size int) *MemCache {
	return &MemCache{Cache: freecache.NewCache(size)}
}

type Operator struct {
	Address string
	// NOTE: for futures iterations
	// Stakes []StakedAmount
}

// type StakedAmount struct{
//   Strategy string
//   Amount string // base10 representation
// }

func genOpKey(opAddr string) []byte {
	return libHash.Keccak256([]byte(opAddr))
}

func (cache *MemCache) CacheOperatorCount(count uint32) error {
	cacheKey := libHash.Keccak256([]byte("skateavs:operator_count"))
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, count)
	if err != nil {
		return err
	}
	return cache.Set(cacheKey, buf.Bytes(), 0)
}

func (cache *MemCache) GetOperatorCount() (*uint32, error) {
	cacheKey := libHash.Keccak256([]byte("skateavs:operator_count"))
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, err // Return zero and the error if unable to get the data
	}

	buf := bytes.NewReader(data)
	var count uint32
	err = binary.Read(buf, binary.BigEndian, &count)
	if err != nil {
		return nil, err // Return zero and the error if unable to decode the data
	}

	return &count, nil // Successfully retrieved and decoded the count
}

func (cache *MemCache) CacheOperator(operator Operator) error {
	key := genOpKey(operator.Address)
	cacheKey := libHash.Keccak256([]byte("skateavs:operator:"), key)
	data, err := json.Marshal(operator)
	if err != nil {
		return err
	}
	return cache.Set(cacheKey, data, 0)
}

func (cache *MemCache) GetOperator(opAddr string) (*Operator, error) {
	key := genOpKey(opAddr)
	cacheKey := libHash.Keccak256([]byte("skateavs:operator:"), key)
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetOperator/Get")
	}
	if len(data) == 0 {
		return nil, nil
	}
	var result Operator
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetOperator/UnMarshal")
	}

	return &result, nil
}

func (cache *MemCache) EvictOperator(opAddr string) bool {
	key := genOpKey(opAddr)
	cacheKey := libHash.Keccak256([]byte("skateavs:operator:"), key)
	return cache.Del(cacheKey)
}

// NOTE: for future uses
func (cache *MemCache) CacheStake(strategy string, value big.Int) error {
	cacheKey := libHash.Keccak256([]byte("skateavs:stake:"), []byte(strategy))
	return cache.Set(cacheKey, value.Bytes(), 0)
}

func (cache *MemCache) GetStake(strategy string) (*big.Int, error) {
	cacheKey := libHash.Keccak256([]byte("skateavs:stake:"), []byte(strategy))
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetStake/Get")
	}
	if len(data) == 0 {
		return nil, nil
	}

	result := new(big.Int)

	return result.SetBytes(data), nil
}
