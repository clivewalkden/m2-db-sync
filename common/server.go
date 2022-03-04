package common

import (
	"fmt"
	"github.com/spf13/viper"
)

var sourceData string

type Server struct {
	Host             string
	Port             int
	User             string
	Group            string
	Domain           []string
	MagentoDir       string
	BackupDir        string
	SharedDir        string
	PhpPath          string
	N98Path          string
	IgnoreTables     []string
	WordPressNetwork bool
}

func ServerSetup(source string) Server {
	sourceData = source
	return Server{
		Host:             returnString("host"),
		Port:             returnInt("port"),
		User:             returnString("user"),
		Group:            returnString("group"),
		Domain:           returnArray("domain"),
		MagentoDir:       returnString("magento_dir"),
		BackupDir:        returnString("backup_dir"),
		SharedDir:        returnString("shared_dir"),
		PhpPath:          returnString("php_path"),
		N98Path:          returnString("n98_path"),
		IgnoreTables:     returnArray("ignore_tables"),
		WordPressNetwork: returnBool("wordpress_network"),
	}
}

func returnString(key string) string {
	if viper.IsSet(fmt.Sprintf("environments.%s.%s", sourceData, key)) {
		return viper.GetString(fmt.Sprintf("environments.%s.%s", sourceData, key))
	}

	return viper.GetString(fmt.Sprintf("%s.%s", "defaults", key))
}

func returnInt(key string) int {
	if viper.IsSet(fmt.Sprintf("environments.%s.%s", sourceData, key)) {
		return viper.GetInt(fmt.Sprintf("environments.%s.%s", sourceData, key))
	}

	return viper.GetInt(fmt.Sprintf("%s.%s", "defaults", key))
}

func returnArray(key string) []string {
	if viper.IsSet(fmt.Sprintf("environments.%s.%s", sourceData, key)) {
		return viper.GetStringSlice(fmt.Sprintf("environments.%s.%s", sourceData, key))
	}

	return viper.GetStringSlice(fmt.Sprintf("%s.%s", "defaults", key))
}

func returnBool(key string) bool {
	if viper.IsSet(fmt.Sprintf("environments.%s.%s", sourceData, key)) {
		return viper.GetBool(fmt.Sprintf("environments.%s.%s", sourceData, key))
	}

	return viper.GetBool(fmt.Sprintf("%s.%s", "defaults", key))
}
