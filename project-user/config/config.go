package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/wushengyouya/project-common/logs"
	"go.uber.org/zap"
)

var AppConf = InitConfig()

type Config struct {
	viper      *viper.Viper
	GrpcConfig *GrpcConfig
	AppConfig  *AppConfig
	EtcdConfig *EtcdConfig
}

type AppConfig struct {
	Addr string
	Name string
}

type GrpcConfig struct {
	Addr    string
	Name    string
	Version string
	Weight  int64
}

type EtcdConfig struct {
	Addrs []string `mapstructure:"addrs"`
}

func (c *Config) ReadGrpcConfig() {
	grpc := new(GrpcConfig)
	err := c.viper.UnmarshalKey("grpc", grpc)
	if err != nil {
		zap.L().Error("grpc config read err: ", zap.Error(err))
	}
	log.Println("grpcconfig: ", grpc)
	c.GrpcConfig = grpc
}

func (c *Config) ReadAppConfig() {
	app := new(AppConfig)
	err := c.viper.UnmarshalKey("app", app)
	if err != nil {
		zap.L().Error("app config read err: ", zap.Error(err))
	}
	log.Println("appconfig: ", app)
	c.AppConfig = app
}

func (c *Config) ReadEtcdConfg() {
	etcdConfig := new(EtcdConfig)
	err := c.viper.UnmarshalKey("etcd", etcdConfig)
	if err != nil {
		zap.L().Error("etcd config read err: ", zap.Error(err))
	}
	log.Println("etcd config : ", len(etcdConfig.Addrs), etcdConfig)
	c.EtcdConfig = etcdConfig
}

// InitConfig 初始化config
func InitConfig() *Config {
	// 配置viper
	viper := viper.New()
	conf := &Config{viper: viper}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	log.Println(workDir + "/config")

	// 读取配置文件
	err := conf.viper.ReadInConfig()
	conf.ReadAppConfig()
	conf.ReadGrpcConfig()
	conf.ReadEtcdConfg()

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
	log.Println("zapconfig: ", lc)
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
