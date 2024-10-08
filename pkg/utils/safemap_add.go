package utils

import "gitea.peekaboo.tech/peekaboo/crushon-backend/pkg/safemap"

func AddItemToArray[K comparable, V any](p *safemap.Map[K, []V], k K, v V) {
	arr, ok := p.Get(k)
	if !ok {
		arr = make([]V, 0)
	}
	arr = append(arr, v)
	p.Set(k, arr)
}
