package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前时间
	now := time.Now()
	fmt.Printf("now: %v\ntype: %T\n", now, now)

	// 时间相关函数
	fmt.Printf("年： %v\n", now.Year())
	fmt.Printf("月：%v\n", now.Month())
	fmt.Printf("月：%v\n", int(now.Month()))
	fmt.Printf("日：%v\n", now.Day())
	fmt.Printf("时：%v\n", now.Hour())
	fmt.Printf("分：%v\n", now.Minute())
	fmt.Printf("秒：%v\n", now.Second())

	//格式话时间日期
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	str := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(str)

	// 使用 now.Format
	str2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(str2)

	// Unix, UnixNano
	// 返回 Unix 时间戳（秒级别）
	fmt.Printf("unix %v\n", now.Unix())
	// 返回 UnixNano 时间戳(纳秒级)
	fmt.Printf("unixnnao %v\n", now.UnixNano())
}
