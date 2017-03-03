package model

import (
	. "common"

)

const (
	CACHE_TIME_OUT = 60
)
func CachePutList(key string,value interface{},timeout int64)  {
	Cache.Put(key,value,timeout)
}

func CacheGetList(key string) interface {}{
	data := Cache.Get(key)
	if data != nil {
		return data
	}else{
		return nil
	}
}