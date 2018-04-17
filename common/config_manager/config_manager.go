package config_manager

import (
	"os"
	"io/ioutil"
	"encoding/xml"
)

type configs struct  {
	ConnectingString string
	LexemeUrl string
	LexemeKey string
	ClientResourceVersion string
	Port string
	Redis string
}

var configData *configs

func getConfig() (config *configs, err error)   {
	if configData != nil {
		return configData, nil
	}
	file, err := os.Open("config.xml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result =  new(configs)
	err = xml.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	configData = result
	return result, nil
}

func GetConfig() *configs {
	config, err := getConfig()
	if err != nil {
		return nil
	} else {
		return config
	}
}

func GetDbConnectionString()( connectionString string, err error){
	config, err := getConfig()
	if err == nil {
		 return config.ConnectingString, nil
	} else {
		return "", err
	}
}

func CheckConfig()  {
	config, err := getConfig()
	if err == nil {
		if config.ConnectingString == "" {
			panic("数据库连接字符串未配置")
		}
		if config.LexemeKey == "" {
			panic("分词系统标识未配置")
		}
		if config.LexemeUrl == "" {
			panic("分词系统URL未配置")
		}
		if config.Port == "" {
			panic("端口号未配置")
		}
		if config.Redis == "" {
			panic("Redis未配置")
		}
	} else {
		panic("配置文件读取失败")
	}
}

