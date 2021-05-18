package cfg

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/toolkits/file"
)

type serverConfig struct {
	MiFalconReportPeriod int
	TracingAddr          string
	EtcdEndPoints        []string
}

type serviceConfig struct {
	AppName       string
	GrpcPort      int
	LogLevel      int
	TimeZone      string
	EnableMonitor bool
	//
	EnableV1WriteBack bool
	EnableV2WriteBack bool
}

type envConfig struct {
	Region   string
	Currency string
}

type DbNameConfig struct {
	MISDb  string
	DealDb string
	QRDb   string
}

type MySqlConfig struct {
	Master       bool
	Host         string
	Port         int
	User         string
	Password     string
	Db           string
	MaxOpenConns int
	MaxIdleSize  int
	ShowSql      bool
}

type redisConfig struct {
	Host               string
	Passwd             string
	PoolSize           int
	DialTimeoutMs      int
	ReadTimeoutMs      int
	WriteTimeoutMs     int
	IdleCheckFrequency int
	IdleTimeoutMs      int
	MaxRetries         int
	ClusterMode        bool
}

type FileServerConfig struct {
	Host       string
	RetryTimes int
	TimeOut    int
	SecretKey  string
	AccessKey  string
}

type globalConfig struct {
	Server     serverConfig
	Service    serviceConfig
	Env        envConfig `toml:"env"`
	DbName     DbNameConfig
	Db         map[string][]MySqlConfig
	Redis      redisConfig
	Includes   []string
	FileServer FileServerConfig
}

var (
	cfg = &globalConfig{}
)

func GetGlobalConfig() *globalConfig {
	return cfg
}

func ParseConfig(cfgPath string) {
	var err error
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("parse config err: %v", err))
		}
	}()

	println(fmt.Sprintf("conf init: confPath=%s", cfgPath))

	if cfgPath == "" {
		err = fmt.Errorf("cfgPath is empty")
		return
	}

	if !file.IsExist(cfgPath) {
		err = fmt.Errorf("conf file is not exists")
		return
	}

	if _, err := toml.DecodeFile(cfgPath, &cfg); err != nil {
		panic(err)
	}

	registerIncludeConfigs()
	for _, includeFile := range cfg.Includes {
		path := includeFilePath(cfgPath, includeFile)
		includeConfig := getIncludeConfig(includeFile)
		if _, err := toml.DecodeFile(path, includeConfig); err != nil {
			panic(err)
		}
	}

	println(fmt.Sprintf("conf init: conf init successful"))
}

func includeFilePath(cfgPath string, includeFileName string) string {
	dir, _ := filepath.Split(cfgPath)
	return filepath.Join(dir, includeFileName)
}
