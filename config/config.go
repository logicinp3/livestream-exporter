package config

import (
    "fmt"

    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)


type Credentials struct {
    AK string `mapstructure:"ak"`
    SK string `mapstructure:"sk"`
}

type Provider struct {
    G04 Credentials `mapstructure:"g04"`
    G13 Credentials `mapstructure:"g13"`
}

type Config struct {
    Haiwei  Provider `mapstructure:"haiwei"`

    Tencent Provider `mapstructure:"tencent"`
}

var AppConfig Config

// Loads configuration from the config file and sets defaults.
func LoadConfig() error {
    // Setting config file
    viper.SetConfigFile("./config/config.yaml")

    // Setting default variable
    viper.SetDefault("default-key", "default-var")

    // Bind environment variable
    //viper.BindEnv("redis.port", "REDIS_PORT")
    viper.AutomaticEnv()

    // Read config file
    if err := viper.ReadInConfig(); err != nil {
        return fmt.Errorf("Error reading config file: %s", err)
    }

    // Unmarshal config file
    if err := viper.Unmarshal(&AppConfig); err != nil {
        return fmt.Errorf("Unable to decode into struct Config: %v", err)
    }

    return nil
}

// Watches for changes and return new context.
func WatchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed: ", e.Name)
        if err := viper.Unmarshal(&AppConfig); err != nil {
            fmt.Printf("Unable to decode into struct Config: %v\n", err)
        }
        //fmt.Println("Updated config:", AppConfig)
    })
}

