package configs

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

// LoadConfig Project runtime config.json
func LoadConfig() (*Config, error) {

	dir := path.Join(path.Dir("config.json"), "./")
	viper.SetConfigType("json")
	viper.SetConfigName("config.json")
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func LoadTestConfig(filepath string) (*Config, error) {
	dir := path.Join(path.Dir("config_test.json"), filepath)
	viper.SetConfigType("json")
	viper.SetConfigName("config_test.json")
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config_test file:", err)
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
