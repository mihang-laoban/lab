package conf

var C Config

type Config struct {
	Base Base
	Db   DbConf
}

type Base struct {
	Port int
	Db []string
	Cache []string
}

type DbConf struct {
	Name string
	Host string
	Port int
	User string
	Password string
	Database string
}