package utils

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// FileCache struct contains data map and mutex for synchronization
type FileCache struct {
	sync.Mutex

	data map[string]*data
	file *os.File
}

type data struct {
	Data    interface{}
	Expired time.Time
}

// NewFileCache creates a new file cache
func NewFileCache(filename string) (*FileCache, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, ErrorWithStack(err)
	}

	cache := &FileCache{
		data: make(map[string]*data),
		file: file,
	}
	if err := cache.loadFromFile(); err != nil {
		return nil, ErrorWithStack(err)
	}
	return cache, nil
}

// Get returns cached value
func (cache *FileCache) Get(key string) interface{} {
	cache.Lock()
	defer cache.Unlock()

	if ret, ok := cache.data[key]; ok {
		if ret.Expired.Before(time.Now()) {
			cache.deleteKey(key)
			return nil
		}
		return ret.Data
	}
	return nil
}

// IsExist checks if value exists in cache
func (cache *FileCache) IsExist(key string) bool {
	cache.Lock()
	defer cache.Unlock()

	if ret, ok := cache.data[key]; ok {
		if ret.Expired.Before(time.Now()) {
			cache.deleteKey(key)
			return false
		}
		return true
	}
	return false
}

// Set caches value with key and expiration time
func (cache *FileCache) Set(key string, val interface{}, timeout time.Duration) error {
	cache.Lock()
	defer cache.Unlock()
	expiration := time.Now().Add(timeout)

	if expiration.Before(time.Now()) {
		return ErrorWithStack(fmt.Errorf("expiration time is already in the past"))
	}

	cache.data[key] = &data{
		Data:    val,
		Expired: expiration,
	}

	return cache.saveToFile()
}

// Delete deletes value from cache
func (cache *FileCache) Delete(key string) error {
	cache.Lock()
	defer cache.Unlock()

	cache.deleteKey(key)
	return cache.saveToFile()
}

// deleteKey deletes key from cache data
func (cache *FileCache) deleteKey(key string) {
	delete(cache.data, key)
}

// loadFromFile loads cache data from file
func (cache *FileCache) loadFromFile() error {
	cache.Lock()
	defer cache.Unlock()

	cache.data = make(map[string]*data)

	decoder := gob.NewDecoder(cache.file)
	if err := decoder.Decode(&cache.data); err != nil && err != io.EOF {
		return ErrorWithStack(err)
	}
	return nil
}

// saveToFile saves cache data to file
func (cache *FileCache) saveToFile() error {
	_, err := cache.file.Seek(0, 0)
	if err != nil {
		return ErrorWithStack(err)
	}
	err = cache.file.Truncate(0)
	if err != nil {
		return ErrorWithStack(err)
	}

	encoder := gob.NewEncoder(cache.file)
	if err := encoder.Encode(cache.data); err != nil {
		return ErrorWithStack(err)
	}
	return nil
}
