package main

import (
	"fmt"
	"log"
	"time"
	"tour/internal/timer"
)

func main() {
	t2()
}

func t2() {

	nowTime := timer.GetNowTime()

	log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
}

func t1() {
	// 测试duration
	res, err := timer.GetCalculateTime(time.Now(), "2h")
	if err != nil {
		log.Println("出错")
	}
	fmt.Println(res)
}
