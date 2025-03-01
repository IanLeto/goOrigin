package utils

import (
	"fmt"
	"strconv"
	"time"
)

// TimeFormat 结构体，用于存储不同格式的时间
type TimeFormat struct {
	Format      string `json:"format"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

func ConvertTime(inputTime string, inputFormat string) (map[string][]TimeFormat, error) {
	var parsedTime time.Time
	var err error

	// 尝试解析 Unix 时间戳（秒或毫秒）
	if timestamp, convErr := strconv.ParseInt(inputTime, 10, 64); convErr == nil {
		if len(inputTime) == 10 { // 秒级时间戳
			parsedTime = time.Unix(timestamp, 0)
		} else if len(inputTime) == 13 { // 毫秒级时间戳
			parsedTime = time.Unix(0, timestamp*int64(time.Millisecond))
		} else {
			return nil, fmt.Errorf("invalid Unix timestamp length")
		}
	} else {
		// 解析标准时间字符串
		parsedTime, err = time.Parse(inputFormat, inputTime)
		if err != nil {
			return nil, fmt.Errorf("invalid input format or time string: %v", err)
		}
	}

	// UTC 和 东八区（北京时间）
	utcTime := parsedTime.UTC()
	cnTime := parsedTime.In(time.FixedZone("CST", 8*3600)) // UTC+8

	// Unix 时间戳
	unixTimestamp := parsedTime.Unix()
	unixMillis := parsedTime.UnixMilli()

	// 存储所有格式的时间
	result := map[string][]TimeFormat{
		"ISO_8601": {
			{"YYYY-MM-DDTHH:mm:ssZ", utcTime.Format("2006-01-02T15:04:05Z"), "标准 UTC 时间"},
			{"YYYY-MM-DDTHH:mm:ss±hh:mm", cnTime.Format("2006-01-02T15:04:05-07:00"), "北京时间（东八区）"},
			{"YYYY-MM-DDTHH:mm:ss.SSSZ", utcTime.Format("2006-01-02T15:04:05.000Z"), "带毫秒的 UTC 时间"},
		},
		"RFC_822_2822": {
			{"EEE, dd MMM yyyy HH:mm:ss Z", cnTime.Format("Mon, 02 Jan 2006 15:04:05 -0700"), "邮件和 HTTP 头格式"},
			{"EEE, dd MMM yyyy HH:mm:ss zzz", cnTime.Format("Mon, 02 Jan 2006 15:04:05 MST"), "带 CST 时区的格式"},
		},
		"Unix_Timestamp": {
			{"Unix Timestamp (seconds)", fmt.Sprintf("%d", unixTimestamp), "Unix 时间戳（秒）"},
			{"Unix Timestamp (milliseconds)", fmt.Sprintf("%d", unixMillis), "Unix 时间戳（毫秒）"},
		},
		"Database_Formats": {
			{"YYYY-MM-DD HH:mm:ss", parsedTime.Format("2006-01-02 15:04:05"), "数据库通用格式"},
			{"YYYY/MM/DD HH:mm:ss", parsedTime.Format("2006/01/02 15:04:05"), "旧系统格式"},
			{"YYYYMMDDHHmmss", parsedTime.Format("20060102150405"), "紧凑格式"},
		},
		"Programming_Languages": {
			{"EEE MMM dd HH:mm:ss zzz yyyy", cnTime.Format("Mon Jan 02 15:04:05 MST 2006"), "Java Date.toString() 格式"},
			{"yyyy-MM-dd'T'HH:mm:ss.SSSXXX", cnTime.Format("2006-01-02T15:04:05.000-07:00"), "Java SimpleDateFormat 格式"},
			{"new Date().toISOString()", utcTime.Format("2006-01-02T15:04:05.000Z"), "JavaScript Date 对象"},
			{"datetime.datetime.now()", parsedTime.Format("2006-01-02 15:04:05.000000"), "Python datetime 格式"},
		},
		"Timezones": {
			{"UTC±hh:mm", utcTime.Format("UTC-07:00"), "通用 UTC 时区格式"},
			{"GMT±hh:mm", utcTime.Format("GMT-07:00"), "格林尼治时间格式"},
			{"IANA Timezone", "Asia/Shanghai", "IANA 时区名称（东八区）"},
			{"IANA Timezone", "America/New_York", "IANA 时区名称（纽约）"},
		},
	}

	return result, nil
}
