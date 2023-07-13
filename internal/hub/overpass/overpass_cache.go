package overpass

import (
	"sync"
	"time"

	gc "github.com/patrickmn/go-cache"
)

var (
	onceStatesCache sync.Once
	statesCache *gc.Cache
)

func initStatesCache() {
	statesCache = gc.New(time.Hour, 15*time.Minute)
}

func StatesCache() *gc.Cache {
	onceStatesCache.Do(initStatesCache)

	return statesCache
}


func CacheSet(city, state string) {
	StatesCache().SetDefault(city, state)
}

func CacheGet(city string) string {
	state, ok := StatesCache().Get(city)
	if !ok {
		return ""
	}

	return state.(string)
}