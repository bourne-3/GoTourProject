package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

// 时间推算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
