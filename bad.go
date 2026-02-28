// package mai

// import "sync"

// var (
// 	mx = sync.Mutex{}
// 	config *Config
// )

// func GetConfig() *Config {
// 	mx.Lock()
// 	defer mx.Unlock()
// 	if config == nil {
// 		config = &Config{}
// 	}
	
// 	return config
// }