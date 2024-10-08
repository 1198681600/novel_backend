package utils

func GetAny[K comparable, V any](m map[K]V, picker func(value V) bool) (V, bool) {
	var v V
	for _, v := range m {
		if picker(v) {
			return v, true
		}
	}
	return v, false
}
