package gokey

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	pairsSet map[string]tuple //contains expiration time and value of a key

	hashFn func([]byte) (string, error)

	sets map[string][]tuple
}

var (
	_ Operations = (*Cache)(nil)

	ErrEmptyKey   = errors.New("key cannot be empty")
	ErrNoExistKey = errors.New("key does not exist")
	ErrExpiredKey = errors.New("key has expired")
)

func newCache() *Cache {
	return &Cache{
		RWMutex:  sync.RWMutex{},
		pairsSet: make(map[string]tuple),
		hashFn:   generateMD5,
	}
}

// Get the values of the key, if this exists in the cache
func (c *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrEmptyKey
	}

	c.RLock()
	defer c.RUnlock()

	keyHashed, err := c.hashFn([]byte(key))
	if err != nil {
		return nil, err
	}

	pair, exists := c.pairsSet[keyHashed]

	if !exists {
		return nil, ErrNoExistKey
	}

	if time.Since(pair.createdAt) > pair.ttl && pair.ttl != -1 {
		delete(c.pairsSet, keyHashed)
		return nil, ErrNoExistKey
	}

	return pair.value, nil
}

// Upsert cache a new key pair or update an existing one
// if ttl is equals to zero the key will not expire
func (c *Cache) Upsert(key string, value []byte, ttl time.Duration) (bool, error) {
	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	c.Lock()
	defer c.Unlock()

	keyHashed, err := c.hashFn([]byte(key))
	if err != nil {
		return false, err
	}

	// redis in generic command:  if (ttl == -1)
	// golang use with functions time.Duration = -1
	c.pairsSet[keyHashed] = tuple{
		ttl:       ttl,
		createdAt: time.Now(),
		value:     value,
	}

	return true, nil
}

func (c *Cache) Delete(key string) (bool, error) {
	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	c.Lock()
	defer c.Unlock()

	keyHashed, err := c.hashFn([]byte(key))
	if err != nil {
		return false, err
	}
	_, exists := c.pairsSet[keyHashed]

	if exists {
		delete(c.pairsSet, keyHashed)
	} else {
		return false, errors.New("key not found")
	}

	return true, nil
}

func (c *Cache) Exists(key string) (bool, error) {
	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	c.RLock()
	defer c.RUnlock()

	keyHashed, err := c.hashFn([]byte(key))
	if err != nil {
		return false, err
	}

	pair, exists := c.pairsSet[keyHashed]

	if !exists {
		return false, ErrNoExistKey
	}

	if time.Since(pair.createdAt) > pair.ttl && pair.ttl != -1 {
		return false, ErrExpiredKey
	}

	return true, nil
}

func (c *Cache) SAdd(key string, vals ...values) (bool, error) {
	hashedKey, err := c.hashFn([]byte(key))
	if err != nil {
		return false, err
	}
	if exKey, ok := c.sets[hashedKey]; !ok {
		for _, val := range vals {
			b, err := json.Marshal(val)
			if err != nil {
				return false, err
			}
			exKey = append(exKey, tuple{
				ttl:       1,
				createdAt: time.Now(),
				value:     b,
			})
		}
		return true, nil
	}

	var tuples []tuple

}

func (c *Cache) SGet() ([]any, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) SDelete() {
	//TODO implement me
	panic("implement me")
}
