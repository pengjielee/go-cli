package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// 定义版本号为2.0.0，表示重大改变
const VERSION = "v2.0.0"

func main() {
	// 输出应用的版本信息
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v", "V"},
		Usage:   "print the version",
	}

	app := &cli.App{
		Name:    "go-cli", //应用名称
		Version: VERSION,  //应用版本
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "date",
				Aliases: []string{"d", "D"},       //别名
				Usage:   "print the current date", //帮助信息
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("date") {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			}
			return nil
		},

		//定义作者信息
		Authors: []*cli.Author{
			{
				Name:  "pengjielee",
				Email: "386276251@qq.com",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
