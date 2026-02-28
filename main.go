package main

import (
	"fmt"
	"sync"
)

type Config struct {
	ConnectionString string
}

func main() {
	for i := 0; i < 100; i++ {
		config = GetConfigWithOnce()
		// config = GetConfig()
		fmt.Println(i, ":", &config)
	}
}

// Bad
var (
	mx     = sync.Mutex{}
	config *Config
)

func GetConfig() *Config {
	if config != nil {
	mx.Lock()
	defer mx.Unlock()
	if config == nil {
		config = &Config{}
		fmt.Println("creating config")
	}
}

	return config
}




//get config with sync.Once

var (
	once sync.Once
	
)

func GetConfigWithOnce() *Config {
	once.Do(func() {
		config = &Config{}
	})
	fmt.Println("get config from old instance")
	return config
}