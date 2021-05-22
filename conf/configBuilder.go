package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigBuilder(confFileName, postfix, confPath string) *viper.Viper {
	vip := viper.New()
	vip.SetConfigName(fmt.Sprintf("%s.%s", confFileName, postfix)) // name of config file (without extension)
	vip.SetConfigType(postfix) // REQUIRED if the config file does not have the extension in the name
	vip.AddConfigPath(confPath)   // path to look for the config file in
	if err := vip.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return vip
}