package cache

import (
	"github.com/tango-contrib/cache"
)

func NewCache()*cache.Caches {
	options := cache.Options{
		//Adapter:       "redis",
		//AdapterConfig: "addr=127.0.0.1:6379,prefix=cache:",
	}

	return cache.New(options)
}
