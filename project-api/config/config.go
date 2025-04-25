package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/wushengyouya/project-common/logs"
	"go.uber.org/zap"
)

var AppConf = InitConfig()

type Config struct {
	viper      *viper.Viper
	AppConfig  *AppConfig
	EtcdConfig *EtcdConfig
}

type AppConfig struct {
	Addr string
	Name string
}

type EtcdConfig struct {
	Addrs []string
}

func (c *Config) ReadAppConfig() {
	app := new(AppConfig)
	err := c.viper.UnmarshalKey("app", app)
	if err != nil {
		zap.L().Error("config read err: ", zap.Error(err))
	}
	c.AppConfig = app
}

// InitConfig 初始化config
func InitConfig() *Config {
	viper := viper.New()
	conf := &Config{viper: viper}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	log.Println(workDir + "/config")
	err := conf.viper.ReadInConfig()
	conf.ReadEtcdConfg()
	conf.ReadAppConfig()
	if err != nil {
		log.Fatalln("InitConfig: ", err)
		return nil
	}
	return conf
}

func (c *Config) ReadEtcdConfg() {
	etcdConfig := new(EtcdConfig)
	err := c.viper.UnmarshalKey("etcd", etcdConfig)
	if err != nil {
		zap.L().Error("etcd config read err: ", zap.Error(err))
	}
	log.Println("etcd config : ", etcdConfig)
	c.EtcdConfig = etcdConfig
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
