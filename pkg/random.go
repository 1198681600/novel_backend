package pkg

import (
	"math/rand"
	"time"
)

func GenerateRandomUpperCaseCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generateRandomCode(charset, length)
}

func GenerateRandomDigitalCode(length int) string {
	const charset = "0123456789"
	return generateRandomCode(charset, length)
}

func GenerateRandomBase62Code(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	return generateRandomCode(charset, length)
}

func generateRandomCode(charset string, length int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}

// RandomInt64FromRange 从范围中随机取一个整数值
func RandomInt64FromRange(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// RandomFloat64FromRange 从范围中随机取一个float64
func RandomFloat64FromRange(min, max float64) float64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Float64()*(max-min) + min
}
