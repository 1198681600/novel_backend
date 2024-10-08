package pkg

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"gitea.peekaboo.tech/peekaboo/crushon-backend/internal/global"
	jsoniter "github.com/json-iterator/go"
	"github.com/lestrrat-go/strftime"
	"github.com/trimmer-io/go-csv"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicInfo(args ...interface{}) {
	info := fmt.Sprintln(args...)
	panic(info)
}

func ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func ToInt32(v string) int32 {
	return int32(ToInt64(v))
}

func ToUint32(v string) uint32 {
	return uint32(ToUint64(v))
}

func ToInt(v string) int {
	return int(ToInt64(v))
}
func ToPointer[T any](t T) *T {
	return &t
}

func ToInt64(v string) int64 {
	if v == "" {
		return 0
	}
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return int64(0)
	}
	return ret
}

func ToUint64(v string) uint64 {
	if v == "" {
		return 0
	}
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0
	}
	return ret
}

func ToFloat64(v string) float64 {
	if v == "" {
		return 0.0
	}
	ret, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0.0
	}
	return ret
}

func ToBool(v string) bool {
	if v == "" {
		return false
	}
	ret, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return ret
}

func BoolToString(v bool) string {
	if v {
		return "true"
	} else {
		return "false"
	}
}

func BoolToUint8(v bool) uint8 {
	if v {
		return 1
	} else {
		return 0
	}
}

func StringToBool(str string) bool {
	switch str {
	case "true":
		return true
	default:
		return false
	}
}

func StringSliceToString(strs []string) string {
	builder := strings.Builder{}
	for idx, str := range strs {
		builder.WriteString(str)
		if idx < len(strs)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func Int32SliceToString(nums []int32) string {
	builder := strings.Builder{}
	for idx, str := range nums {
		builder.WriteString(Int32ToString(str))
		if idx < len(nums)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func Int32ToString(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func StringToInt32(str string) (num int32) {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return
	}
	num = int32(value)
	return
}

// UnionInt64 合并去重
func UnionInt64(a, b []int64) []int64 {
	set := make(map[int64]struct{}, len(a)+len(b))

	storeInt64ToMap(set, a)
	storeInt64ToMap(set, b)

	ret := make([]int64, 0, len(set))

	return storeMapToInt64(ret, set)
}

// IntersectInt64 取交集
func IntersectInt64(a, b []int64) []int64 {
	setA := make(map[int64]struct{}, len(a))
	setB := make(map[int64]struct{}, len(a))

	storeInt64ToMap(setA, a)
	storeInt64ToMap(setB, b)

	for val := range setA {
		if _, ok := setB[val]; !ok {
			delete(setA, val)
		}
	}
	ret := make([]int64, 0, len(setA))
	return storeMapToInt64(ret, setA)
}

func JsonToStringArr(sor string) (result []string) {
	err := json.Unmarshal([]byte(sor), &result)
	if err != nil {
		return nil
	}
	return
}

func StringArrToJson(sor []string) (result string) {
	marshal, err := json.Marshal(sor)
	if err != nil {
		return "{}"
	}
	return string(marshal)
}

func storeInt64ToMap(m map[int64]struct{}, a []int64) {
	for _, val := range a {
		m[val] = struct{}{}
	}
}

func storeMapToInt64(a []int64, m map[int64]struct{}) []int64 {
	for val := range m {
		a = append(a, val)
	}
	return a
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func MinFloat(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func Bytes2Int64(b []byte) (ret int64, err error) {
	if len(b) == 0 {
		return
	}

	var flag = int64(1)

	if b[0] == '-' {
		flag = -1
	}

	if b[0] == '+' || b[0] == '-' {
		b = b[1:]
	}

	for _, v := range b {
		if v < '0' || v > '9' {
			err = errors.New("invalid number")
			break
		}
		ret = ret*10 + int64(v-'0')
	}

	ret *= flag
	return
}

// MaxMinInt64 依次返回两个数的 较大值 与 较小值
func MaxMinInt64(x, y int64) (int64, int64) {
	if x > y {
		return x, y
	}
	return y, x
}

func MaxMinUint64(x, y uint64) (uint64, uint64) {
	if x > y {
		return x, y
	}
	return y, x
}

// ShuffleIntSlice 用于打乱int切片的顺序
func ShuffleIntSlice(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

func Uint64ToBytes(x uint64) []byte {
	var (
		number [64]byte
		cnt    = 63
	)

	if x == 0 {
		return []byte{'0'}
	}

	for x > 0 {
		number[cnt] = byte(x%10) + '0'
		x = x / 10
		cnt--
	}
	return number[cnt+1:]
}

// GetDayZeroTimeStamp 获取当天零点时间戳 注意时区
func GetDayZeroTimeStamp() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return tm.Unix()
}

// GetUTCDayZeroTimeStamp 获取当天零点时间戳
func GetUTCDayZeroTimeStamp() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return tm.Unix()
}

// GetHourZeroTimeStamp 获取这个小时的零点时间戳
func GetHourZeroTimeStamp() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	return tm.Unix()
}

// GetWeek 获取当天周数 注意时区
func GetWeek(timeStamp int64) (int64, error) {
	s, err := strftime.New("%Y%U")
	if err != nil {
		return 0, err
	}
	tm := time.Unix(timeStamp, 0).Local()
	//get the first day of week by Sunday
	weekDay := tm.Weekday()
	switch weekDay {
	case time.Monday:
		tm = time.Unix(timeStamp-86400, 0).Local()
	case time.Tuesday:
		tm = time.Unix(timeStamp-2*86400, 0).Local()
	case time.Wednesday:
		tm = time.Unix(timeStamp-3*86400, 0).Local()
	case time.Thursday:
		tm = time.Unix(timeStamp-4*86400, 0).Local()
	case time.Friday:
		tm = time.Unix(timeStamp-5*86400, 0).Local()
	case time.Saturday:
		tm = time.Unix(timeStamp-6*86400, 0).Local()
	}
	return ToInt64(s.FormatString(tm)), nil
}

func ContentJsonToStruct(content string, res interface{}) (err error) {
	err = json.Unmarshal([]byte(content), res)
	if err != nil {
		global.Logger.Error("ContentJsonToStruct failed to unmarshal", zap.Error(err))
		return
	}

	return
}

func JsonToStruct(path string, res interface{}) (err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		global.Logger.Error("failed to read file", zap.String("path", path), zap.Error(err))
		return
	}

	err = json.Unmarshal(content, res)
	if err != nil {
		global.Logger.Error("failed to unmarshal", zap.Error(err))
		return
	}

	return
}

// ParseTimeV1 2023-07-11T11:05:11.835Z to in64
func ParseTimeV1(timeStr string) (res int64, err error) {
	timestamp, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return
	}
	return timestamp.Unix(), nil
}

func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}

func GetLock(key string, ttl time.Duration) (ok bool) {
	// 抢占锁
	ok, err := global.RedisClient.SetNX(context.Background(), key, "lock", ttl).Result()
	if err != nil {
		global.Logger.Error("failed to get lock", zap.Error(err))
		return
	}
	if !ok {
		return
	}
	return
}

func DelLock(key string) (err error) {
	err = global.RedisClient.Del(context.Background(), key).Err()
	if err != nil {
		global.Logger.Error("failed to del lock", zap.Error(err))
		return
	}
	return
}

func GetTimestampOfMidnightInUTC8() time.Time {
	utcTime := time.Now().UTC()                // 获取当前的UTC时间
	utcTimePlus8 := utcTime.Add(8 * time.Hour) // 加上8小时得到东八区时间

	// 创建东八区的时区
	utc8Zone := time.FixedZone("UTC+8", 8*60*60)

	// 构造当天0点时间，并应用东八区时区
	midnight := time.Date(utcTimePlus8.Year(), utcTimePlus8.Month(), utcTimePlus8.Day(), 0, 0, 0, 0, utc8Zone)

	return midnight
}

var ErrTimeout = fmt.Errorf("timeout")

// RunWithTimeout 执行指定的函数 fn，并在超时时取消执行。
func RunWithTimeout[V any](fn func() (V, error), timeout time.Duration) (value V, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct {
		value V
		err   error
	}, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recover Panic from SearchCharactersV2 subscriber. %+v\n", r)
				// 打印堆栈信息
				buf := make([]byte, 4096)
				runtime.Stack(buf, false)
				fmt.Printf("Stack Trace:\n%s\n", buf)
			}
		}()
		value, err := fn()
		done <- struct {
			value V
			err   error
		}{value: value, err: err}
	}()

	select {
	case result := <-done:
		return result.value, result.err
	case <-ctx.Done():
		err = ErrTimeout
		return
	}
}

func RunWithTimeoutV2[V any](fn func(context.Context) (V, error), timeout time.Duration) (value V, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct {
		value V
		err   error
	}, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recover Panic from SearchCharactersV2 subscriber. %+v\n", r)
				// 打印堆栈信息
				buf := make([]byte, 4096)
				runtime.Stack(buf, false)
				fmt.Printf("Stack Trace:\n%s\n", buf)
			}
		}()
		value, err := fn(ctx)
		done <- struct {
			value V
			err   error
		}{value: value, err: err}
	}()

	select {
	case result := <-done:
		return result.value, result.err
	case <-ctx.Done():
		err = ErrTimeout
		return
	}
}

func ContainsChinese(text string) bool {
	// 使用正则表达式匹配中文字符
	regex := regexp.MustCompile("\\p{Han}")
	return regex.MatchString(text)
}

// GetCurrentSubscriptionMonthPeriodStart 针对年付用户，获取当前的订阅周期开始时间
func GetCurrentSubscriptionMonthPeriodStart(currentPeriodStart time.Time) time.Time {
	now := time.Now()
	result := currentPeriodStart
	for result.Before(now) {
		result = result.AddDate(0, 1, 0)
	}
	result = result.AddDate(0, -1, 0)
	return result
}

// GetUTCMonthStart 获取UTC时间本月的开始时间戳
func GetUTCMonthStart() time.Time {
	now := time.Now().UTC()
	year, month, _ := now.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return startOfMonth
}

func StringToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func ContainStringsV2(content string, findWords []string) (ok bool) {
	content = strings.ToLower(content)
	for _, v := range findWords {
		// 构建正则表达式模式，匹配关键词前后不是字母或数字的情况
		pattern := regexp.QuoteMeta(v)
		re, err := regexp.Compile(pattern)
		if err != nil {
			global.Logger.Error("regexp compile error", zap.Error(err))
			continue
		}
		// 在文本中查找匹配项
		matches := re.FindAllString(content, -1)
		if len(matches) > 0 {
			ok = true
			return
		}
	}
	return
}

// ContainStrings 判断文本中是否包含词汇列表 原文本会转成小写匹配 且前后不是字母或数字的情况
func ContainStrings(content string, findWords []string) (ok bool) {
	content = strings.ToLower(content)
	for _, v := range findWords {
		// 构建正则表达式模式，匹配关键词前后不是字母或数字的情况
		pattern := `\b` + regexp.QuoteMeta(v) + `\b`
		re, err := regexp.Compile(pattern)
		if err != nil {
			global.Logger.Error("regexp compile error", zap.Error(err))
			continue
		}
		// 在文本中查找匹配项
		matches := re.FindAllString(content, -1)
		if len(matches) > 0 {
			ok = true
			return
		}
	}
	return
}

// GetFirstContainString 获取str中第一个在findWords中的词
func GetFirstContainString(content string, findWords []string) (result string) {
	content = strings.ToLower(content)
	for _, v := range findWords {
		// 构建正则表达式模式，匹配关键词前后不是字母或数字的情况
		pattern := `\b` + regexp.QuoteMeta(v) + `\b`
		re, err := regexp.Compile(pattern)
		if err != nil {
			global.Logger.Error("regexp compile error", zap.Error(err))
			continue
		}
		// 在文本中查找匹配项
		matches := re.FindAllString(content, -1)
		if len(matches) > 0 {
			result = v
			return
		}
	}
	return
}

// SliceBounds 泛型函数，接受一个切片、开始位置和结束位置作为参数，并返回一个切片
func SliceBounds[T any](s []T, start int, end int) []T {
	// 检查开始位置和结束位置是否在切片范围内
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start > end {
		return nil
	}

	// 提取切片中的数据
	return s[start:end]
}

func FindKeyWords(str string, keyWords []string) (hitKeyWords []string) {
	lowStr := strings.ToLower(str)
	for _, v := range keyWords {
		// 构建正则表达式模式，匹配关键词前后不是字母或数字的情况
		pattern := `(?i)(?U:\b|\W|^)` + regexp.QuoteMeta(v) + `(?U:\b|\W|$)`

		re, err := regexp.Compile(pattern)
		if err != nil {
			global.Logger.Error("regexp compile error", zap.Error(err))
			continue
		}
		// 在文本中查找匹配项
		matches := re.FindAllString(lowStr, -1)
		if len(matches) > 0 {
			hitKeyWords = append(hitKeyWords, v)
		}
	}
	return
}

func ArrayToMap[T comparable](arr []T) (rlt map[T]interface{}) {
	rlt = make(map[T]interface{})
	for _, a := range arr {
		rlt[a] = struct{}{}
	}
	return
}

// Deduplication 重复数据删除
func Deduplication[T comparable](arr []T) (rlt []T) {
	rlt = make([]T, 0)
	m := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			rlt = append(rlt, v)
			m[v] = struct{}{}
		}
	}
	return
}

func ContainsAny[S ~[]E, E comparable](s S, items []E) bool {
	return slices.ContainsFunc(s, func(item E) bool {
		return slices.Contains(items, item)
	})
}

func SubStringLimitLenByRune(source string, limit int) (target string) {
	runeTitle := []rune(source)
	if len(runeTitle) > limit {
		runeTitle = runeTitle[:limit]
	}
	target = string(runeTitle)
	return
}

// RandomSample 从一个切片中随机取出 n 个元素
func RandomSample[T any](slice []T, n int) []T {
	if n > len(slice) {
		return slice
	}

	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(slice))
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = slice[perm[i]]
	}
	return result
}

func PInt64(i int64) *int64 {
	return &i
}

func PBool(i bool) *bool {
	return &i
}

func PString(s string) *string {
	return &s
}

func CSVToStruct[T any](data []byte) (result []T, err error) {
	result = make([]T, 0)
	if err = csv.Unmarshal(data, &result); err != nil {
		return
	}
	return
}

func CsvReadStream[T any](r *os.File) error {
	dec := csv.NewDecoder(r)

	// read and decode the file header
	line, err := dec.ReadLine()
	if err != nil {
		return err
	}
	if _, err = dec.DecodeHeader(line); err != nil {
		return err
	}
	// loop until EOF (i.e. dec.ReadLine returns an empty line and nil error);
	// any other error during read will result in a non-nil error
	for {
		// read the next line from stream
		line, err = dec.ReadLine()

		// check for read errors other than EOF
		if err != nil {
			return err
		}

		// check for EOF condition
		if line == "" {
			break
		}

		// decode the record
		var v = new(T)
		if err = dec.DecodeRecord(v, line); err != nil {
			return err
		}

		// process the record here
		//Process(v)
	}
	return nil
}

func GetRankingString(ranking int64) (res string) {
	if ranking%100 >= 11 && ranking%100 <= 13 {
		res = strconv.FormatInt(ranking, 10) + "th"
	} else {
		switch ranking % 10 {
		case 1:
			res = strconv.FormatInt(ranking, 10) + "st"
		case 2:
			res = strconv.FormatInt(ranking, 10) + "nd"
		case 3:
			res = strconv.FormatInt(ranking, 10) + "rd"
		default:
			res = strconv.FormatInt(ranking, 10) + "th"
		}
	}
	return
}
