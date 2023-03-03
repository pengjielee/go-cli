package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

const VERSION = "1.0.2"

func getIP() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return fmt.Sprint("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return fmt.Sprint(ipnet.IP.String())
					}
				}
			}
		}
	}
	return ""
}

func main() {
	// var showVersion bool
	// var showDate bool
	// var showIP bool
	var showVersion, showDate, showIP bool

	flag.BoolVar(&showVersion, "v", false, "查看版本")
	flag.BoolVar(&showDate, "d", false, "查看日期")
	flag.BoolVar(&showIP, "ip", false, "查看本地IP")

	flag.Parse()

	if showVersion {
		fmt.Println(VERSION)
	}

	if showDate {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	}

	if showIP {
		fmt.Println(getIP())
	}
}
