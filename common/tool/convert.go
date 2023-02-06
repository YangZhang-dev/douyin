package tool

import "time"

// UnixToTime 时间戳转时间
func UnixToTime(timeUnix int64, layout string) (time.Time, error) {
	local, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return time.Time{}, err
	}
	return tt, nil
}
