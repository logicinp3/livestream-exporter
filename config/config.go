package config

import (
    "fmt"

    "github.com/spf13/viper"
)

func init() {
    // Setting config file
    viper.SetConfigFile("./config/config.yaml")

    // Read config to viper
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    // Get config
    a := viper.Get("haiwei")
    fmt.Println(a)

    k := viper.Get()
    fmt.Println(k)
}