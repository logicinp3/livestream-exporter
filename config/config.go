package config

import (
    "fmt"

    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

type HwAuthConfig struct {
    HUAWEICLOUD_SDK_AK string  `mapstructure:"ak"`
    HUAWEICLOUD_SDK_SK string  `mapstructure:"sk"`
}

type TcAuthConfig struct {
    TENCENTCLOUD_SDK_AK string  `mapstructure:"ak"`
    TENCENTCLOUD_SDK_SK string  `mapstructure:"sk"`
}

type Config struct {
    Haiwei struct {
        G04 HwAuthConfig  `mapstructure:"g04"`
        G13 HwAuthConfig  `mapstructure:"g13"`
    } `mapstructure:haiwei`

    Tencent struct {
        G04 TcAuthConfig  `mapstructure:"g04"`
        G13 TcAuthConfig  `mapstructure:"g13"`
    } `mapstructure:tencent`
}

var AppConfig Config

//func init() {
func LoadConfig() {
    // Setting config file
    viper.SetConfigFile("./config/config.yaml")

    // Read config to viper
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }


    // Get config
    //a := viper.Get("haiwei.0.g04.ak")
    //fmt.Println(a)
    return viper.Unmarshal(&AppConfig)
}

func init() {
    // Dynamic watch config
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotiff

    viper.SetDefault("default-key", "default-var")
    viper.AutomaticEnv()
}
