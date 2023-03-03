package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Now struct {
	Temp    string `json:"temp"`
	Text    string `json:"text"`
	WindDir string `json:"windDir"`
}
type Result struct {
	Code string `json:"code"`
	Now  Now
}

// 定义版本号为2.0.0，表示重大改变
const VERSION = "v2.0.0"

func getWeatherInfo(cityCode string) string {
	cityMap := map[string]string{
		"bj": "101010100",
		"sh": "101020100",
		"tj": "101030100",
		"hz": "101210101",
		"cd": "101270101",
		"zz": "101180101",
	}
	var cityName string

	switch cityCode {
	case "bj":
		cityName = "北京"
	case "tj":
		cityName = "天津"
	case "sh":
		cityName = "上海"
	case "cd":
		cityName = "成都"
	case "hz":
		cityName = "杭州"
	case "zz":
		cityName = "郑州"
	default:
		cityName = ""
	}

	if value, ok := cityMap[cityCode]; !ok {
		return "该城市暂不支持"
	} else {
		url := fmt.Sprintf("https://devapi.heweather.net/v7/weather/now?key=5a193793d8854a0395a890426d0e6b62&location=%s", value)
		resp, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "Error status code"
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Sprintf("Error io readall: %s", err.Error())
		}
		var res Result
		err = json.Unmarshal(body, &res)
		if err != nil {
			return fmt.Sprintf("Error json unmarshal: %s", err.Error())
		}
		if res.Code != "200" {
			return fmt.Sprintf("Error Code: %s", "请求失败")
		}
		return fmt.Sprintf("%s：%s，%s摄氏度，%s", cityName, res.Now.Text, res.Now.Temp, res.Now.WindDir)
	}
}

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

		Commands: []*cli.Command{
			{
				Name:    "weather",
				Aliases: []string{"w"},
				Usage:   "查询天气，支持（bj,tj,sh,hz,cd,zz）",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "city code",   // 命令名称
						Aliases: []string{"c"}, // 简写
						Value:   "bj",          // 默认值
					},
				},
				Action: func(c *cli.Context) error {
					cityName := c.String("city")
					fmt.Println(getWeatherInfo(cityName))
					return nil
				},
			},
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
