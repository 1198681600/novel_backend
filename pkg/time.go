package pkg

import (
	"gorm.io/datatypes"
	"time"
)

const (
	defaultDateStr = "2000-01-01"
)

var (
	DefaultDate     datatypes.Date
	DefaultDateTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
)

func init() {
	if value, err := time.Parse("2006-01-02", defaultDateStr); err != nil {
		panic("failed to parse default date")
	} else {
		DefaultDate = datatypes.Date(value)
	}
}

const BaseStartTimestamp = int64(1577836800000)

func ConvertToDay(current, baseTimestamp int64) int64 {
	return ConvertToHour(current, baseTimestamp) / 24
}

func ConvertToHour(current, baseTimestamp int64) int64 {
	return (current - baseTimestamp) / int64(time.Millisecond/1000*60*60)
}

func GetUTCStartOfDay(inputTime time.Time) time.Time {
	// Get the year, month, and day components of the input time
	year, month, day := inputTime.Date()

	// Create a new time.Time representing the start of the day (00:00:00)
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return startOfDay
}

func GetDurationFromNowToEndOfDay() time.Duration {
	// 获取当前时间
	now := time.Now()
	// 获取今天的结束时间
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	// 计算剩余时间
	timeRemaining := endOfDay.Sub(now)
	return timeRemaining / 1000
}

func GetDurationFromNowToEndOfDaySeconds() time.Duration {
	// 获取当前时间
	now := time.Now()
	// 获取今天的结束时间
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	// 计算剩余时间
	timeRemaining := endOfDay.Sub(now)
	return timeRemaining / time.Second
}

func GetDurationFromNowToEndOfDayNormal() time.Duration {
	// 获取当前时间
	now := time.Now()
	// 获取今天的结束时间
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	// 计算剩余时间
	timeRemaining := endOfDay.Sub(now)
	return timeRemaining
}

// GetDayOfYear 一年中的第几天
func GetDayOfYear() (year, day int) {
	now := time.Now()
	dayOfYear := now.YearDay()
	return now.Year(), dayOfYear
}

// GetWeekOfYear 一年中的第几周
func GetWeekOfYear() (year, week int) {
	now := time.Now()
	return now.ISOWeek()
}

// GetMonthOfYear 一年中的第几月
func GetMonthOfYear() (year, month int) {
	now := time.Now()
	return now.Year(), int(now.Month())
}

// GetQuarterOfYear 一年中的第几季
func GetQuarterOfYear() (year, quarter int) {
	// 获取当前时间
	now := time.Now()
	// 获取当前月份
	month := now.Month()
	// 计算当前季度
	quarter = (int(month)-1)/3 + 1
	return now.Year(), quarter
}

var startDate = time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

func GetDayFrom2023() int32 {
	endDate := time.Now()
	return int32(endDate.Sub(startDate).Hours() / 24)
}

func GetWeekFrom2023() int32 {
	days := GetDayFrom2023()
	return days / 7
}

func GetMonthFrom2023() int32 {
	endDate := time.Now()
	return int32(endDate.Year()-startDate.Year())*12 + int32(endDate.Month())
}

func GetQuarterFrom2023() int32 {
	months := GetMonthFrom2023()
	if months%3 == 0 {
		return months / 3
	}
	return (months / 3) + 1
}

var BaseTimeLayout = "2006-01-02 15:04:05"

func ParseFormatTimeToTime(timeStr string) *time.Time {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"

	// 解析字符串为 time.Time 类型
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil
	}
	return &t
}

func ConvertTimeToFormatTime(layout string, time2 time.Time) string {
	// 定义时间格式
	return time2.Format(layout)
}
