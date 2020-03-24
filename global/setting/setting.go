package setting

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type WEB struct {
	URL  string `toml:"url"`
	Port string `toml:"port"`
}

type Mongo struct {
	URI      string `toml:"uri"`
	Database string `toml:"database"`
}

type Log struct {
	Database   string `toml:"database"`
	Collection string `toml:"collection"`
}

// Config 所有的配配置信息结构体
type _config struct {
	Title  string `toml:"title"`
	Secret string `toml:"secret"`
	WEB    WEB    `toml:"web"`
	Mongo  Mongo  `toml:"mongo"`
	Log    Log    `toml:"log"`
}

var Config _config

// 根据环境变量解析对应的配置信息
func InitConfig(env string) {
	// 读取配置文件
	file, err := ioutil.ReadFile("./config/" + env + ".toml")
	if err != nil {
		log.Fatal("读取配置文件失败", err.Error())
	}
	// 解析配置文件
	if _, err = toml.Decode(string(file), &Config); err != nil {
		log.Fatal("解析配置文件失败", err.Error())
	}
}
