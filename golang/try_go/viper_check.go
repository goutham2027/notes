package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Name   string
		Nested NestedConfig
	}

	NestedConfig struct {
		Key int
	}
)

func GetConfig() (Config, error) {
	var c Config
	err := ReadAppConfig(&c, []string{
		"",
	})
	return c, err
}

func ReadAppConfig(cfg *Config, paths []string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalExact(&cfg); err != nil {
		return err
	}
	fmt.Println(cfg)

	return nil
}

func GetENV() {
	c, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(c.Nested.Key)
	fmt.Printf("%+v\n", c)
}
