package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/blinkbean/dingtalk"
)

const (
	serviceName           string = "Flight-Go"
	currentServiceVersion string = "v0.1.0"
	logMaxAge                    = time.Hour * 24
	logRotationTime              = time.Hour * 24
	logPath               string = ""
	logFileName           string = "Flight-Go.log"
)

var cli *dingtalk.DingTalk

type Config struct {
	DingTalkAppKey    string `json:"dingTalkAppKey"`
	DingTalkAppSecret string `json:"dingTalkAppSecret"`
}

func loadConfig() (*Config, error) {
	filePath := "config.json" // Replace with the actual file path
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
func main() {

	config, err := loadConfig()
	if err != nil {
		// Handle error
	}

	cli = dingtalk.InitDingTalkWithSecret(config.DingTalkAppKey, config.DingTalkAppSecret)
	logger = NewLogger(logMaxAge, logRotationTime, logPath, logFileName)
	// 命令行初始化
	commandLineInit()
	commands := flightCommands
	for {
		args := os.Args
		if len(args) > 1 {
			for _, cmd := range commands {
				if cmd.Run != nil && cmd.Name() == args[1] {
					err := cmd.Flag.Parse(args[2:])
					if err != nil {
						os.Exit(1)
					}
					args = cmd.Flag.Args()
					if len(args) > 0 {
						// 初始化数据
						initCityNameCodeData()
						os.Exit(cmd.Run(args))
					}
					logger.Errorf("[Flight-Go]命令参数错误!")
					break
				}
			}
			commandUsage()
		} else {
			logger.Errorf("[Flight-Go]命令参数错误!")
			commandUsage()
		}
		time.Sleep(time.Hour * 2) // 每两小时执行一次
	}
}
