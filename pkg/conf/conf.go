package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Cfg *Config

type Config struct {
	Gin struct {
		Mode   string `mapstructure:"mode"`
		Listen string `mapstructure:"listen"`
	} `mapstructure:"gin"`

	Mail struct {
		From     string `mapstructure:"from"`
		UserName string `mapstructure:"username"`
		AuthCode string `mapstructure:"auth_code"`
	} `mapstructure:"mail"`

	RPC struct {
		MainnetEth     string `mapstructure:"mainnet_eth"`
		MainnetPolygon string `mapstructure:"mainnet_polygon"`
		TestnetGoerli  string `mapstructure:"testnet_goerli"`
		TestnetWallaby string `mapstructure:"testnet_wallaby"`
		TestnetWMumbai string `mapstructure:"testnet_mumbai"`
	} `mapstructure:"rpc"`

	Database struct {
		Dialect     string `mapstructure:"dialect"`
		Host        string `mapstructure:"host"`
		Port        uint   `mapstructure:"port"`
		Dbname      string `mapstructure:"dbname"`
		User        string `mapstructure:"user"`
		Password    string `mapstructure:"password"`
		MaxIdleConn int    `mapstructure:"max_idle_conn"`
		MaxOpenConn int    `mapstructure:"max_open_conn"`
		Debug       bool   `mapstructure:"debug"`
		AutoMigrate bool   `mapstructure:"auto_migrate"`
		SSLMode     string `mapstructure:"ssl_mode"`
	} `mapstructure:"database"`

	Redis struct {
		Network  string `mapstructure:"network"`
		Host     string `mapstructure:"host"`
		Port     uint   `mapstructure:"port"`
		Db       int    `mapstructure:"db"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	Web3Storage struct {
		Token string `mapstructure:"token"`
	} `mapstructure:"web_storage"`

	Session struct {
		ExpiresTime int64  `mapstructure:"expires_time"`
		KeyPairs    string `mapstructure:"key_pairs"`
	} `mapstructure:"session"`
}

func LoadConfig(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	Cfg = cfg
	fmt.Println(Cfg)
}
