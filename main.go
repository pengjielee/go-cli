package main

import (
	"flag"
	"fmt"
	"time"
)

const VERSION = "1.0.1"

func main() {
	var showVersion bool
	var showDate bool
	flag.BoolVar(&showVersion, "v", false, "查看版本")
	flag.BoolVar(&showDate, "d", false, "查看日期")

	flag.Parse()

	if showVersion {
		fmt.Println(VERSION)
	}

	if showDate {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	}
}
