package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/wushengyouya/project-common/logs"
)

var AppConf = InitConfig()

type Config struct {
	viper *viper.Viper
}

// InitConfig 初始化config
func InitConfig() *Config {
	viper := viper.New()
	conf := &Config{viper: viper}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "\\config")
	log.Println(workDir + "\\config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln("InitConfig: ", err)
		return nil
	}
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{}
	err := c.viper.UnmarshalKey("zap", lc)
	if err != nil {
		log.Fatalln("InitZapLog: ", err)
	}

	err = logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) InitRedisOptions() *redis.Options {
	redisOptions := &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
	return redisOptions
}
