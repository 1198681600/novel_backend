package pkg

import (
	"gitea.peekaboo.tech/peekaboo/crushon-backend/internal/global"
	"sync"
	"time"

	"go.uber.org/zap"
)

type cacheCell struct {
	ts      int64
	content interface{}
}
type simpleCache struct {
	sync.RWMutex
	cache     map[interface{}]*cacheCell
	ctime     int64 // 缓存时长
	cSize     int   // 缓存数量
	cCleanGap int   // 定时清理过期缓存,定时间隔
	isDel     bool  // 定时清理是否有删除元素
}

// NewSimpleCache cleanGap:the frequency of check and clean cache
// cleanGap <= 0 means the cache should lasting,so don't clean it
func NewSimpleCache(cacheTime int64, cacheSize, cleanGap int) *simpleCache {
	cache := &simpleCache{ctime: cacheTime, cSize: cacheSize, cCleanGap: cleanGap,
		cache: make(map[interface{}]*cacheCell, cacheSize)}
	if cache.cCleanGap > 0 {
		go cache.loopClean()
	}
	return cache
}

func (r *simpleCache) SetCache(key, content interface{}) {
	r.Lock()
	defer r.Unlock()

	cell := &cacheCell{}
	cell.ts = time.Now().Unix()
	cell.content = content
	r.cache[key] = cell
}

func (r *simpleCache) GetCache(key interface{}) (interface{}, bool) {
	r.RLock()
	defer r.RUnlock()

	cell, ok := r.cache[key]
	if !ok {
		return 0, false
	}
	return cell.content, true
}

func (r *simpleCache) loopClean() {
	for {
		time.Sleep(time.Duration(r.cCleanGap) * time.Second)

		start := time.Now()
		r.Lock()
		oldLen := len(r.cache)
		for k, v := range r.cache {
			if start.Unix()-v.ts > r.ctime {
				delete(r.cache, k)
				r.isDel = true
			}
		}
		newLen := len(r.cache)
		r.Unlock()

		duration := time.Since(start)
		if r.isDel && duration > 1*time.Second {
			global.Logger.Error("simpleCache loopClean", zap.Any("timeCost", time.Since(start)), zap.Int("oldLen", oldLen), zap.Int("newLen", newLen))
		}
		r.isDel = false
	}
}
