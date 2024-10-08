package utils

import "io"

func Reader2String(resp io.Reader) (s string) {
	bytes, err := io.ReadAll(resp)
	if err != nil {
		s = ""
	}
	s = string(bytes)
	return s
}
