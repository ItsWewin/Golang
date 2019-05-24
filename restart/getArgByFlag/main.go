package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		user string
		pwd  string
		host string
		port string
	)
	// &user 接受命令行 -u 后面的参数
	// "u", 就是 -u 制定参数
	// "root", 默认值
	// "user name, default value root" 参数说明
	flag.StringVar(&user, "u", "root", "user name, default value 'root'")
	flag.StringVar(&pwd, "p", "123456", "password, default value '123456'")
	flag.StringVar(&host, "h", "localhost", "host, default value 'localhost'")
	flag.StringVar(&port, "P", "3306", "port, default value '3306'")

	flag.Parse()

	fmt.Printf("user = %v, password = %v, host = %v, port = %v", user, pwd, host, port)
}
