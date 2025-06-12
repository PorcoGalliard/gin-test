package configs

import (
	"github.com/spf13/viper"
)

var config *Config

type option struct {
	ConfigFolders []string
	ConfigFile string
	ConfigType string
}

func Init(opts ...Option) error {
	opt := &option{
		ConfigFolders: getDefaultConfigFolders(),
		ConfigFile: getDefaultConfigFile(),
		ConfigType: getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, configFolder := range opt.ConfigFolders {
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(opt.ConfigFile)
	viper.SetConfigType(opt.ConfigType)
	viper.AutomaticEnv()

	config = new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(&config)

}

type Option func(*option)

func getDefaultConfigFolders() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.ConfigFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.ConfigFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.ConfigType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}
